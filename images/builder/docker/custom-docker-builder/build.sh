#!/bin/bash
set -o pipefail
IFS=$'\n\t'
GIT_SSH=
GIT_CONFIG=

DOCKER_SOCKET=/var/run/docker.sock

if [ ! -e "${DOCKER_SOCKET}" ]; then
    echo "Docker socket missing at ${DOCKER_SOCKET}"
    exit 1
fi

if [ -d "${SOURCE_SECRET_PATH}" ]; then
    echo "Source secret mounted to ${SOURCE_SECRET_PATH}"
    USERNAME="${SOURCE_SECRET_PATH}/username"
    PASSWORD="${SOURCE_SECRET_PATH}/password"
    SSHPRIVKEY="${SOURCE_SECRET_PATH}/ssh-privatekey"

    if [ -f "${SSHPRIVKEY}" ]; then
        #echo "ssh-privatekey found..."
        cp ${SSHPRIVKEY} /tmp/ssh-privatekey
        export SSHPRIVKEY=/tmp/ssh-privatekey
        chmod 0600 ${SSHPRIVKEY}
        export GIT_SSH="/tmp/git.sh"
        echo -e '#!/bin/sh\nssh -i ${SSHPRIVKEY} -o StrictHostKeyChecking=false $@' > "${GIT_SSH}"
        chmod +x "${GIT_SSH}"
    else
        if [ -f "${PASSWORD}" ] ; then
            #echo "password found..."
            PASSWORD=$(cat ${PASSWORD})
            if [ -f "${USERNAME}" ];then
                #echo "username found..."
                USERNAME=$(cat ${USERNAME})
            else
                #echo "username not found..."
                USERNAME="builder"
            fi
            GITCREDENTIAL="${HOME}/.git-credentials"
            export GIT_CONFIG="${HOME}/.gitconfig"
            echo "https://${USERNAME}:${PASSWORD}@github.com" > "${GITCREDENTIAL}"
            echo -e '[credential]\n\thelper = store'  > "${GIT_CONFIG}"
        fi
    fi
fi

if [ -n "${OUTPUT_IMAGE}" ]; then
    TAG="${OUTPUT_REGISTRY}/${OUTPUT_IMAGE}"
fi

if [[ "${SOURCE_REPOSITORY}" != "ssh://"* ]] && [[ "${SOURCE_REPOSITORY}" != "git://"* ]] && [[ "${SOURCE_REPOSITORY}" != "git@"* ]]; then
    URL="${SOURCE_REPOSITORY}"
    if [[ "${URL}" != "http://"* ]] && [[ "${URL}" != "https://"* ]]; then
        URL="https://${URL}"
    fi
    echo "Checking source url ${URL}"
    git ls-remote $URL --heads > /dev/null
    if [ $? != 0 ]; then
        echo "Could not access source url: ${SOURCE_REPOSITORY}"
        exit 1
    fi
fi

function show_git_repo_info {
    git submodule update --init --recursive
    echo "ref: $(git rev-parse --abbrev-ref HEAD)"
    echo "commitid: $(git rev-parse --verify HEAD)"
    echo "author name: $(git --no-pager show -s --format=%an HEAD)"
    echo "author email: $(git --no-pager show -s --format=%ae HEAD)"
    echo "committer name: $(git --no-pager show -s --format=%cn HEAD)"
    echo "committer email $(git --no-pager show -s --format=%ce HEAD)"
    echo "date: $(git --no-pager show -s --format=%ad HEAD)"
    echo "message: $(git --no-pager show -s --format=%s HEAD)"
}

if [ -n "${SOURCE_REF}" ]; then
    BUILD_DIR=$(mktemp --directory -t build.XXXXXXXXX)
    git clone --recursive "${SOURCE_REPOSITORY}" "${BUILD_DIR}"
    if [ $? != 0 ]; then
        echo "Error trying to fetch git source: ${SOURCE_REPOSITORY}"
        exit 1
    fi
    pushd "${BUILD_DIR}"
    echo "Checkout branch ${SOURCE_REF}"
    git checkout "${SOURCE_REF}" > /dev/null
    if [ $? != 0 ]; then
        echo "Error trying to checkout branch: ${SOURCE_REF}"
        exit 1
    fi
    TAGCOMMITID="${OUTPUT_REGISTRY}/$(echo $OUTPUT_IMAGE | awk -F ':' '{print $1}'):${SOURCE_REF}-$(git rev-parse --verify --short HEAD)"
    show_git_repo_info
    popd
    echo "Build output to ${TAG} ${TAGCOMMITID}"
    docker build --rm -t "${TAG}" "${BUILD_DIR}"
    docker tag -f "${TAG}" "${TAGCOMMITID}"
else
    docker build --rm -t "${TAG}" "${SOURCE_REPOSITORY}"
fi

if [[ -d /var/run/secrets/openshift.io/push ]] && [[ ! -e /root/.dockercfg ]]; then
    cp /var/run/secrets/openshift.io/push/.dockercfg /root/.dockercfg
fi

if [ -n "${OUTPUT_IMAGE}" ] || [ -s "/root/.dockercfg" ]; then
    docker push "${TAG}"
    if [ $? != 0 ]; then
        echo "Push ${TAG} failed."
    else
        echo "Push ${TAG} successfully."
    fi
    if [ -n "${TAGCOMMITID}" ]; then
        docker push "${TAGCOMMITID}"
        if [ $? != 0 ]; then
            echo "Push ${TAGCOMMITID} failed."
        else
            echo "Push ${TAGCOMMITID} successfully."
        fi
    fi
fi

# Erik Petrosyan ©
RELEASE_VERSION=1.0.0
ROLLBACK_VERSION=1.0.0

GIT_REPO=git@github.com:PetrosyanDev/shtem-api.git
GIT_BLD_BRANCH=main
DEV_HOST=erik@shtemaran.am
PRD_HOST=erik@shtemaran.am
DEV_BASE=/home/erik
PRD_BASE=/home/erik
SSH_PORT=22
DEPLOY_DIR=shtem-api
# CERTS_DIR:=$(DEV_BASE)/imedcs-tls
# REGISTRY=stg.aestportal.com:19090
IMAGE=shtem-api

## Clonning Repo to Stage Server
init:
	ssh ${DEV_HOST} -p ${SSH_PORT} "git clone ${GIT_REPO}"
	ssh ${DEV_HOST} -p ${SSH_PORT} "cd ${DEPLOY_DIR} && git checkout ${GIT_BLD_BRANCH}"

## Keep Sources UpToDate
pull:
	ssh ${DEV_HOST} -p ${SSH_PORT} "cd ${DEPLOY_DIR} && git checkout ${GIT_BLD_BRANCH} && git pull origin ${GIT_BLD_BRANCH}"

# ## Generating Protobuff Files
# proto:
# 	@mkdir -p sources/pkg
# 	protoc -I=imedcs-idls/imedcs-api --go_out=sources/pkg --go-grpc_out=sources/pkg imedcs-idls/imedcs-api/api.proto
# 	protoc -I=imedcs-idls/imedcs-storage --go_out=sources/pkg --go-grpc_out=sources/pkg imedcs-idls/imedcs-storage/storage.proto
# 	protoc -I=imedcs-idls/imedcs-emails --go_out=sources/pkg --go-grpc_out=sources/pkg imedcs-idls/imedcs-emails/emails.proto

## Installing all Dependencies on Local Machine
deps:
	go mod tidy -compat=1.21.3
	go mod vendor

## Running Tests on Local Machine
test:
	go test -cover shtem-api/...

## Running on Local Machine
run: 
	go run shtem-api/... --cfg secrets/local.json

# ## Running on Local Machine with TLS
# run-tls: test
# 	@NONSENCE=${NONSENCE} go run imedcs/... --tls --cfg secrets/local.json

build: test pull
	mkdir -p build/api
	cd sources/cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../build/api/app
	scp -P ${SSH_PORT} -r build ${DEV_HOST}:${DEV_BASE}/${DEPLOY_DIR}/
	ssh ${DEV_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} docker-compose -f ${DEPLOY_DIR}/docker/build.yml build"
# 	ssh ${DEV_HOST} -p ${SSH_PORT} "docker tag ${IMAGE}:${RELEASE_VERSION} ${REGISTRY}/${IMAGE}:${RELEASE_VERSION}"
# 	ssh ${DEV_HOST} -p ${SSH_PORT} "docker push ${REGISTRY}/${IMAGE}:${RELEASE_VERSION}"
	@echo "BUILT IMAGE: ${IMAGE}:${RELEASE_VERSION}"

build-prd: test
	mkdir -p build/api
	cd sources/cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../build/api/app
	scp -P ${SSH_PORT} -r docker ${DEV_HOST}:${DEV_BASE}/${DEPLOY_DIR}/
	scp -P ${SSH_PORT} -r build ${DEV_HOST}:${DEV_BASE}/${DEPLOY_DIR}/
	ssh ${DEV_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} docker-compose -f ${DEPLOY_DIR}/docker/build.yml build"
	@echo "BUILT IMAGE: ${IMAGE}:${RELEASE_VERSION}"

## Building and Deploying on Staging
deploy-dev: build
	scp -P ${SSH_PORT} secrets/dev.json ${DEV_HOST}:${DEV_BASE}/${DEPLOY_DIR}/secrets.json
	ssh ${DEV_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} DIR=${DEV_BASE}/${DEPLOY_DIR} docker stack deploy -c ${DEPLOY_DIR}/docker/run.yml erik --with-registry-auth"
	ssh ${DEV_HOST} -p ${SSH_PORT} "rm -f ${DEPLOY_DIR}/secrets.json"
	@echo "DEPLOYED on STAGING! VERSION is: ${RELEASE_VERSION}"

## Building and Deploying on Production
deploy-prd: build-prd
	ssh ${PRD_HOST} -p ${SSH_PORT} "mkdir -p ${DEPLOY_DIR}/docker"
	scp -P ${SSH_PORT} -r docker/run.yml ${PRD_HOST}:${PRD_BASE}/${DEPLOY_DIR}/docker/
	scp -P ${SSH_PORT} secrets/prd.json ${PRD_HOST}:${PRD_BASE}/${DEPLOY_DIR}/secrets.json
	ssh ${PRD_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} DIR=${PRD_BASE}/${DEPLOY_DIR} MODE=release NONS=${NONSENCE} docker stack deploy -c ${DEPLOY_DIR}/docker/run.yml imedcs --with-registry-auth"
	ssh ${PRD_HOST} -p ${SSH_PORT} "rm -f ${DEPLOY_DIR}/secrets.json"
	@echo "DEPLOYED on PRODUCTION! VERSION is: ${RELEASE_VERSION}"

## Rolling Back on Production by one deploy
revert:
	ssh ${PRD_HOST} -p ${SSH_PORT} "docker pull ${REGISTRY}/${IMAGE}:${ROLLBACK_VERSION}"
	ssh ${PRD_HOST} -p ${SSH_PORT} "docker service update --image ${REGISTRY}/${IMAGE}:${ROLLBACK_VERSION} --force imedcs_${IMAGE} --with-registry-auth"
	@echo "ROLLED BACK on PRODUCTION! now VERSION is: ${ROLLBACK_VERSION}"

## Purge Docker Caches on Build Server
cleanup:
	ssh ${DEV_HOST} -p ${SSH_PORT} "docker system prune -f"
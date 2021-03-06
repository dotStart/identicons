APPLICATION_VERSION := 0.1.0
APPLICATION_COMMIT_HASH := `git log -1 --pretty=format:"%H"`
APPLICATION_TIMESTAMP := `date --utc "+%s"`

LDFLAGS :=-X 'github.com/dotstart/identicons/internal/build.version=${APPLICATION_VERSION}' -X 'github.com/dotstart/identicons/internal/build.commitHash=${APPLICATION_COMMIT_HASH}' -X 'github.com/dotstart/identicons/internal/build.timestampRaw=${APPLICATION_TIMESTAMP}'

DOCKER := $(shell command -v docker 2> /dev/null)
GO := $(shell command -v go 2> /dev/null)
TAR := $(shell command -v tar 2> /dev/null)
export

PLATFORMS := darwin/amd64 linux/amd64 linux/arm linux/arm64 windows/amd64/.exe

# magical formula:
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
ext = $(word 3, $(temp))

all: $(PLATFORMS)

check-env:
	@echo "==> Checking prerequisites"
	@echo -n "Checking for go ... "
ifndef GO
	@echo "Not Found"
	$(error "go is unavailable")
endif
	@echo $(GO)
	@echo -n "Checking for tar ... "
ifndef TAR
	@echo "Not Found"
	$(error "tar is unavailable")
endif
	@echo $(TAR)
	@echo ""

clean:
	@echo "==> Clearing previous build data"
	@rm -rf target/ || true
	@rm -rf build/package/licenses/ || true
	@rm -rf build/package/identicons_* || true
	@$(GO) clean -cache

licenses:
	@echo "==> collecting 3rd party licenses"
	@$(GO) mod download
	@$(GO) install github.com/google/go-licenses@latest
	@go-licenses save --force "github.com/dotstart/identicons/cmd/identicons" --save_path="target/licenses/" || true

.ONESHELL:
$(PLATFORMS): check-env
	@export GOOS=$(os);
	@export GOARCH=$(arch);

	@echo "==> Building ${os}-${arch}"
	@$(GO) build -v -ldflags "${LDFLAGS}" -o target/$(os)-$(arch)/identicons$(ext) github.com/dotstart/identicons/cmd/identicons

	@echo "==> Creating ${os}-${arch}.tar.gz"
	@$(TAR) -C "target/$(os)-$(arch)/" -czvf "target/identicons_$(os)-$(arch).tar.gz" "identicons$(ext)"

test: check-env
	@echo "==> running tests"
	@$(GO) test ./...

docker-prepare: licenses
	@echo "==> preparing docker build environment"
	@echo -n "Checking for docker ... "
ifndef DOCKER
	@echo "Not Found"
	$(error "docker is unavailable")
endif
	@echo $(DOCKER)
	@echo ""

	@cp -r target/licenses/ build/package/
	@cp LICENSE build/package/licenses/
	@cp target/linux-amd64/identicons build/package/identicons_amd64
	@cp target/linux-arm/identicons build/package/identicons_arm
	@cp target/linux-arm64/identicons build/package/identicons_arm64

docker: docker-prepare
	@echo "==> building docker container"
	@docker build -t ghcr.io/dotstart/identicons:${APPLICATION_VERSION} build/package/

docker-multiarch-versioned: docker-prepare
	@echo "==> building versioned multiarch docker container"
	@docker buildx build --push -t ghcr.io/dotstart/identicons:${APPLICATION_VERSION} -t dotstart/identicons:${APPLICATION_VERSION} --platform linux/amd64,linux/arm64,linux/arm/v7 build/package/

docker-multiarch-latest: docker-prepare
	@echo "==> building latest multiarch docker container"
	@docker buildx build --push -t ghcr.io/dotstart/identicons:latest -t dotstart/identicons:latest --platform linux/amd64,linux/arm64,linux/arm/v7 build/package/

docker-multiarch-dev: docker-prepare
	@echo "==> building dev multiarch docker container"
	@docker buildx build --push -t ghcr.io/dotstart/identicons:dev -t dotstart/identicons:dev --platform linux/amd64,linux/arm64,linux/arm/v7 build/package/

.PHONY: all

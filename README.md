# Template repository for Go projects

# TODOS

## You Really Should Do These

- [ ] Replace occurences of `gitlab.com/zephinzer/template-go-package` with the path of your own package
- [ ] Ensure [LICENSE](./LICENSE) contains the intended license
- [ ] Update the `CMD_NAME` variable in the [`Makefile`](./Makefile)
- [ ] Update the `MAINTAINER` variable in the [`Makefile`](./Makefile)
- [ ] Update the `DOCKER_IMAGE_PATH` variable in the [`Makefile`](./Makefile)

## For Some Use Cases

- [ ] If you're intending to apply versioning via Git tags, set the `DEPLOY_KEY` variable in the Gitlab CI/CD variables (run `make .keys/versioning` to generate your keys, set `DEPLOY_KEY` to the value of the file at `./keys/versioning/id_rsa.base64` and add a Deploy Key to your project using the value of the file at `./keys/versioning/id_rsa.pub`)
- [ ] If you're publishing a Docker image, set the `DOCKER_REGISTRY_USERNAME` and `DOCKER_REGISTRY_PASSWORD` in the Gitlab CI/CD variables
- [ ] If you're using Code Climate, set the `CC_TEST_REPORTER_ID` in the Gitlab CI/CD variables
- [ ] If you're publishing a binary to Github Releases via Travis, set the `RELEASE_TOKEN` in the Travis CI/CD variables
- [ ] If you're syncing this repository from Gitlab to Github, remember to set your sync settings to push to the Github repository so Travis can take over and release the application

# Contributing

## Development Overview

### Directory Layout

This repository follows the "Standard Go Project Layout" as laid out at [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout) where possible/sensible.

### Makefile Recipes

| Recipe | Description |
| --- | --- |
| `deps` | Installs dependencies |
| `run` | Runs the application |
| `test` | Runs tests |
| `install` | Installs the application locally |
| `build` | Build the application |
| `build_production` | Build the application with symbol stripping without using the cache |
| `build_static` | Build the application with static linking |
| `build_static_production` | Build the application with static linking and symbol stripping without using the cache |
| `compress` | Compresses the built application using `upx` |
| `image` | Build the Docker image for the application |
| `image_test` | Runs `container-structure-test` tests on the built image |
| `image_export` | Saves the Docker image as a tarball at `./build/image.tar.gz` |
| `image_import` | Loads the Docker image from a tarball at `./build/image.tar.gz` |
| `.keys/versioning` | Creates a set of keys usable for version bumping |

## Process

1. Clone this repository
2. Make your changes on `master`
3. Issue a Pull Request/Merge Request
4. Await a review
5. Merge on approval

# License

This project is licensed under [the MIT license](./LICENSE).

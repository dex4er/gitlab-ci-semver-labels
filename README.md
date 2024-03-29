# gitlab-ci-semver-labels

[![GitHub](https://img.shields.io/github/v/tag/dex4er/gitlab-ci-semver-labels?label=GitHub)](https://github.com/dex4er/gitlab-ci-semver-labels)
[![Snapshot](https://github.com/dex4er/gitlab-ci-semver-labels/actions/workflows/snapshot.yaml/badge.svg)](https://github.com/dex4er/gitlab-ci-semver-labels/actions/workflows/snapshot.yaml)
[![Release](https://github.com/dex4er/gitlab-ci-semver-labels/actions/workflows/release.yaml/badge.svg)](https://github.com/dex4er/gitlab-ci-semver-labels/actions/workflows/release.yaml)
[![Trunk Check](https://github.com/dex4er/gitlab-ci-semver-labels/actions/workflows/trunk.yaml/badge.svg)](https://github.com/dex4er/gitlab-ci-semver-labels/actions/workflows/trunk.yaml)
[![Docker Image Version](https://img.shields.io/docker/v/dex4er/gitlab-ci-semver-labels?label=Docker&logo=docker&sort=semver)](https://hub.docker.com/r/dex4er/gitlab-ci-semver-labels)
[![Amazon ECR Image Version](https://img.shields.io/docker/v/dex4er/gitlab-ci-semver-labels?label=Amazon%20ECR&logo=Amazon+AWS&sort=semver)](https://gallery.ecr.aws/dex4er/gitlab-ci-semver-labels)

Bump the semver for a Gitlab CI project based on merge request labels.

If no `current` command was used nor any of `bump` command then labels of the
Merge Requests are taken either from `$CI_MERGE_REQUEST_LABELS` environment
variable or from details of the Merge Requests pointed by the commit message
from `$CI_COMMIT_MESSAGE` environment variable.

Commit message should contain the string `See merge request PROJECT!NUMBER`.

To fetch the details of the Merge Request the tool needs the Gitlab API token
with the `read_api` scope. The token is taken from the `$GITLAB_TOKEN`
environment variable by default.

Versions printed by this tool are normalized. It means that `v` prefix is
always trimmed from the output.

The best result is when merge trains are enabled in the merge options for the
project. In this case, it is possible to verify the bumped version before the
actual merge is done.

## Usage

```sh
gitlab-ci-semver-labels [command] [flags]
```

### Docker

From DockerHub:

```sh
docker run -it dex4er/gitlab-ci-semver-labels [command] [flags]
```

or from Amazon ECR Public:

```sh
docker run -it public.ecr.aws/q8i3x1g6/gitlab-ci-semver-labels [command] [flags]
```

Supported tags:

- vX.Y.Z-linux-amd64
- vX.Y.Z-linux-arm64
- vX.Y.Z
- vX.Y
- vX
- latest

### Available Commands

```console
  bump        Bump version
  current     Show current version
  help        Help about any command
```

#### bump

```console
  initial     Set to initial version without checking labels
  major       Bump major version without checking labels
  minor       Bump minor version without checking labels
  patch       Bump patch version without checking labels
```

### Flags

```console
      --commit-message-regexp REGEXP     REGEXP for commit message after merged MR (default "(?s)(?:^|\\n)See merge request (?:\\w[\\w.+/-]*)?!(\\d+)")
  -d, --dotenv-file FILE                 write dotenv format to FILE
  -D, --dotenv-var NAME                  variable NAME in dotenv file (default "VERSION")
  -f, --fail                             fail if merge request are not matched
  -T, --fetch-tags                       fetch tags from git repo (default true)
  -t, --gitlab-token-env VAR             name for environment VAR with Gitlab token (default "GITLAB_TOKEN")
  -g, --gitlab-url URL                   URL of the Gitlab instance (default "https://gitlab.com")
  -h, --help                             help for gitlab-ci-semver-labels
      --initial-label-regexp REGEXP      REGEXP for initial release label (default "(?i)initial.release|semver(.|::)initial")
  -V  --initial-version VERSION          initial VERSION for initial release (default "0.0.0")
      --major-label-regexp REGEXP        REGEXP for major (breaking) release label (default "(?i)(major|breaking).release|semver(.|::)(major|breaking)")
      --minor-label-regexp REGEXP        REGEXP for minor (feature) release label (default "(?i)(minor|feature).release|semver(.|::)(minor|feature)")
      --patch-label-regexp REGEXP        REGEXP for patch (fix) release label (default "(?i)(patch|fix).release|semver(.|::)(patch|fix)")
      --prerelease-label-regexp REGEXP   REGEXP for prerelease label (default "(?i)pre.?release")
  -p, --project PROJECT                  PROJECT id or name (default $CI_PROJECT_ID)
  -P, --prerelease                       bump version as prerelease
  -r, --remote-name NAME                 NAME of git remote (default "origin")
  -v, --version                          VERSION for gitlab-ci-semver-labels
  -C, --work-tree DIR                    DIR to be used for git operations (default ".")
```

### Configuration

Some options can be read from the configuration file
`.gitlab-ci-semver-labels.yml`:

```yaml
commit-message-regexp: (?s)(?:^|\n)See merge request (?:\w[\w.+/-]*)?!(\d+)
dotenv-file: ""
dotenv-var: VERSION
fail: false
fetch-tags: true
gitlab-token-env: GITLAB_TOKEN
gitlab-url: https://gitlab.com
initial-label-regexp: (?i)initial.release|semver(.|::)initial
initial-version: 0.0.0
major-label-regexp: (?i)(major|breaking).release|semver(.|::)(major|breaking)
minor-label-regexp: (?i)(minor|feature).release|semver(.|::)(minor|feature)
patch-label-regexp: (?i)(patch|fix).release|semver(.|::)(patch|fix)
prerelease-label-regexp: (?i)pre.?release
project: dex4er/gitlab-ci-semver-labels
remote-name: origin
work-tree: .
```

### Environment variables

Any option might be overridden with an environment variable with the name the
same as an option with the prefix `GITLAB_CI_SEMVER_LABELS_` and an option name
with all capital letters with a dash character replaced with an underscore. Ie.:

Additionally, `$CI_PROJECT_ID` is used as a default `project` option value and
`$CI_SERVER_URL` as a default `gitlab-url` option value.

The `GITLAB_CI_SEMVER_LABELS_LOG` environment variable changes log level for messages
generated by this tool: `TRACE`, `DEBUG`, `WARNING` or `ERROR` (the default).

## CI

Example `.gitlab-ci.yml`:

```yaml
variables:
  DOCKER_IO: docker.io

stages:
  - semver
  - release
  - label

semver:validate:
  stage: semver
  rules:
    - if: $CI_MERGE_REQUEST_LABELS =~ /semver::/ && $CI_MERGE_REQUEST_EVENT_TYPE == 'merge_train'
  image:
    name: $DOCKER_IO/dex4er/gitlab-ci-semver-labels
    entrypoint: [""]
  variables:
    GIT_DEPTH: 0
  script:
    - gitlab-ci-semver-labels current || true
    - gitlab-ci-semver-labels bump --fail
  cache: []

semver:bump:
  stage: semver
  rules:
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH && $CI_COMMIT_MESSAGE =~ /(^|\n)See merge request (\w[\w.+\/-]*)?!\d+/s
  image:
    name: $DOCKER_IO/dex4er/gitlab-ci-semver-labels
    entrypoint: [""]
  variables:
    GIT_DEPTH: 0
  script:
    - gitlab-ci-semver-labels current || true
    - gitlab-ci-semver-labels bump --dotenv-file=semver.env
  artifacts:
    reports:
      dotenv: semver.env
  cache: []

release:
  stage: release
  needs:
    - semver:bump
  rules:
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH && $CI_COMMIT_MESSAGE =~ /(^|\n)See merge request (\w[\w.+\/-]*)?!\d+/s
  image: registry.gitlab.com/gitlab-org/release-cli
  script:
    - if [ -n "$VERSION" ]; then
      release-cli create
      --name "v$VERSION"
      --description "Automatic release by gitlab-ci-semver-labels"
      --tag-name "v$VERSION"
      --ref $CI_COMMIT_SHA;
      else
      echo "No new version. Release skipped.";
      fi
  cache: []

label:
  stage: label
  needs:
    - release
  rules:
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH && $CI_COMMIT_MESSAGE =~ /(^|\n)See merge request (\w[\w.+\/-]*)?!\d+/s
  image:
    name: registry.gitlab.com/gitlab-org/cli:v1.35.0
    entrypoint: [""]
  variables:
    GITLAB_HOST: $CI_SERVER_URL
  script:
    - MR=$(echo "$CI_COMMIT_MESSAGE" | sed -n '/^See merge request [A-Za-z0-9.+\/-]*![0-9][0-9]*$/s/^See merge request [A-Za-z0-9.+\/-]*!\([0-9][0-9]*\)$/\1/p');
      if [ -n "$VERSION" ] && [ -n "$MR" ]; then
      glab api -X PUT "projects/:id/merge_requests/$MR" --field "add_labels=v::${VERSION#v}" --silent;
      else
      echo "No new version. Label skipped.";
      fi
  cache: []
```

# You can find the new timestamped tags here: https://hub.docker.com/r/gitpod/workspace-full/tags
FROM gitpod/workspace-full:2022-11-06-15-53-10

# Generally try to avoid using sudo in a Dockerfile, we make an exception as it is recommended for GitPod setup
# https://www.gitpod.io/docs/configure/workspaces/workspace-image
RUN curl -O https://downloads.1password.com/linux/debian/amd64/stable/1password-cli-amd64-latest.deb && sudo dpkg -i 1password-cli-amd64-latest.deb

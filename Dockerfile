# Starting from golang image
FROM golang:latest as go-docker

# Our final image should have ros:foxy already installed
FROM ros:foxy

# Get the binaries from previous docker stage:
COPY --from=go-docker /usr/local/go /usr/local/go

# Install extra packages
RUN apt update
RUN apt install -y ros-foxy-rosidl-runtime-c ros-foxy-rosidl-generator-c vim

# Create the main user:
# ARG user=rclgo
# RUN useradd -ms /bin/bash ${user}
# USER ${user}
# WORKDIR /home/${user}
# RUN  mkdir -p /home/${user}/go/bin && \
#      mkdir -p /home/${user}/go/src

# Environment variables
# ENV PATH $PATH:/usr/local/go/bin
# ENV GOPATH /home/${user}/go
# ENV GOROOT /usr/local/go

# Checkout rclgo code:
#RUN git clone https://github.com/juaruipav/rclgo /home/${user}/go/src/rclgo
# COPY . /home/${user}/go/src/rclgo

RUN  mkdir -p /go/bin && \
     mkdir -p /go/src

ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /go
COPY . /go/src/rclgo
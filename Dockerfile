# Builder stage
FROM ubuntu:18.04 AS builder

# Install build dependencies
RUN apt-get update &&\
  apt-get install -y --no-install-recommends\
  build-essential \
  libgd-dev \
  libzzip-dev \
  libopencv-highgui-dev \
  cmake \
  wget \
  protobuf-compiler \
  libprotobuf-dev \
  libopencv-videoio-dev \
  libjsoncpp-dev \
  software-properties-common &&\
  add-apt-repository ppa:longsleep/golang-backports && \
  apt-get update &&\
  apt-get install -y --no-install-recommends\
  golang-go &&\
  rm -rf /var/lib/apt/lists/*

# Download and build YARP
RUN cd /tmp && \
  wget https://github.com/robotology/yarp/archive/v2.3.72.tar.gz && \
  tar xzvf v2.3.72.tar.gz && \
  mkdir /yarp && \
  cd /yarp && \
  cmake -DSKIP_ACE=TRUE /tmp/yarp-* && \
  make

# Build makesweet
COPY ./makesweet/src /makesweet/src
COPY ./makesweet/CMakeLists.txt /makesweet/CMakeLists.txt
RUN cd /makesweet && \
  mkdir build && \
  cd build && \
  cmake -DUSE_OPENCV=ON -DUSE_DETAIL=ON -DYARP_DIR=/yarp .. && \
  make VERBOSE=1

# Create reanimator script
RUN echo "#!/bin/bash" > /reanimator && \
  echo "/makesweet/build/bin/reanimator \"\$@\"" >> /reanimator && \
  chmod u+x /reanimator

# Build server
COPY ./server /makesweetServer
RUN cd /makesweetServer && \
  go mod tidy && \
  go build -o /makesweetServer/start .

# Final stage
FROM ubuntu:18.04

# Install runtime dependencies
RUN apt-get update &&\
  apt-get install -y --no-install-recommends \
  libgd-dev \
  libzzip-dev\
  libopencv-highgui-dev \
  libjsoncpp-dev \
  protobuf-compiler \
  libprotobuf-dev \
  libopencv-videoio-dev &&\
  apt-get clean &&\
  rm -rf /var/lib/apt/lists/*

# Copy built files from the builder stage
COPY --from=builder /yarp/ /yarp/
COPY --from=builder /makesweet/build/ /makesweet/build/
COPY --from=builder /reanimator /reanimator
COPY --from=builder /makesweetServer/start /makesweetServer/start

# Copy templates
COPY ./makesweet/templates/ /makesweet/templates/

# Set entrypoint
ENTRYPOINT ["/makesweetServer/start"]

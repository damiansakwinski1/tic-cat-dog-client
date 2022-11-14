FROM golang:1.19.3-buster

RUN apt-get update -y && apt-get install -y \
    gcc \
    libc6-dev \
    libglu1-mesa-dev \
    libgl1-mesa-dev \
    libxcursor-dev \
    libxi-dev \
    libxinerama-dev \
    libxrandr-dev \
    libxxf86vm-dev \
    libasound2-dev \
    pkg-config

WORKDIR /app
COPY . .
RUN go build
RUN chmod +x tic-cat-dog-client
CMD tic-cat-dog-client
FROM golang:1.16

RUN apt update && apt install -y \
    git gpg gnupg gpg-agent socat \
    ca-certificates \
    zsh \
    curl \
    wget \
    fonts-powerline \
    procps \
    build-essential \
    librdkafka-dev -y

RUN groupadd -g 2000 go \
&& useradd -m -u 2001 -g go guipalm4

USER guipalm4    

RUN mkdir -p /home/guipalm4/go/src

WORKDIR /go/src 
ENV PATH="/go/bin:${PATH}"


RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.3/zsh-in-docker.sh)" -- \
    -t amuse \
    -p git \
    -p https://github.com/zdharma-continuum/fast-syntax-highlighting  \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions \
    -a 'export TERM=xterm-256color'
    
RUN echo 'HISTFILE=/home/go/zsh/.zsh_history' >> ~/.zshrc 

CMD ["tail", "-f", "/dev/null"]
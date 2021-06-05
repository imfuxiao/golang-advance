FROM ubuntu:20.04

ARG GOVERSION
ENV GOVERSION=${GOVERSION}

# 时区设置
ENV TZ="Asia/Shanghai"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list

RUN apt-get update \
	&& apt-get install -y binutils vim gdb git exuberant-ctags gcc tmux build-essential cmake python3-dev ripgrep curl \
	&& rm -rf /var/lib/apt/lists/*

RUN curl -L -O https://golang.google.cn/dl/go${GOVERSION}.linux-amd64.tar.gz \
	&& rm -rf /usr/local/go  && tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz && rm -rf *.tar.gz

ENV GOPATH=$HOME/go GOROOT=/usr/local/go GOPROXY="https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct" FZF_DEFAULT_OPTS='--height 40% --layout=reverse --border'

ENV PATH="$PATH:$GOROOT/bin:$GOPATH/bin"

RUN go get -u github.com/go-delve/delve/cmd/dlv \
	&& go get -u github.com/jstemmer/gotags \
	&& go get -u github.com/sqs/goreturns \
	&& go get -u github.com/junegunn/fzf

RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
	https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

COPY .vimrc /root/.vimrc

# ignore vim install plug exit 1
RUN vim -E -s -u /root/.vimrc +PlugInstall +qall && echo 'success' || echo 'success'

#Cài đặt ubuntu và các gói cần thiết
FROM ubuntu

RUN apt update
RUN apt install -y easy-rsa
RUN apt install -y curl
RUN rm -rf /var/lib/apt/lists/*

#Cài đặt golang
ENV GOLANG_VERSION 1.17.2
RUN curl -sSL https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz \
    | tar -v -C /usr/local -xz
ENV PATH /usr/local/go/bin:$PATH
RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

#Lấy những file cần thiết
COPY server.go /go/src/server.go
RUN mkdir /go/src/Uploads

#Config easy-rsa
RUN mkdir /easy-rsa
RUN ln -s /usr/share/easy-rsa/* /easy-rsa/
RUN chmod 700 /easy-rsa
COPY vars /easy-rsa/vars
COPY sign.sh /go/src/sign.sh
WORKDIR /go/src
RUN /easy-rsa/easyrsa init-pki
RUN echo "CA" | /easy-rsa/easyrsa build-ca nopass


EXPOSE 8080
CMD ["go", "run", "server.go"]
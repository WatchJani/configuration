FROM ubuntu:20.04

ENV GO_VERSION 1.23.4
ENV GO_BINARY go${GO_VERSION}.linux-amd64.tar.gz
ENV GO_URL https://go.dev/dl/${GO_BINARY}
ENV GO_INSTALL_DIR /usr/local

RUN apt-get update && apt-get install -y \
    wget \
    tar \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN wget ${GO_URL} -P /tmp && \
    tar -C $GO_INSTALL_DIR -xvzf /tmp/${GO_BINARY} && \
    rm /tmp/${GO_BINARY}

ENV PATH $GO_INSTALL_DIR/go/bin:$PATH

RUN go version


# Postavljanje radnog direktorija unutar kontejnera
WORKDIR /app

# Kopiranje lokalnog Go koda u kontejner
COPY . .

# Instalacija ovisnosti (ako ih ima) i izgradnja Go aplikacije
RUN go mod tidy
RUN go build -o myapp ./cmd

# Pokretanje aplikacije kada kontejner starta
CMD ["./myapp"]


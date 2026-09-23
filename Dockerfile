# Estágio 1
# Definindo versão minimalista do Linux
FROM golang:1.20-alpine AS builder

# Instalação do Git para poder baixar dependência
RUN apk add --no-cache git

# Diretório de trabalho dentro do container onde os comandos serão executados
WORKDIR /app

# Copia os arquivos de definição de dependências
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia todos os arquivos do diretório local para o container
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o tangle-hornet-api main.go


# Estágio 2
FROM alpine:latest

# INstala certificados para conexão HTTPS e o bash
RUN apk --no-cache add ca-certificates bash

# Diretório de trabalho do container final
WORKDIR /root/

# Traz o binário copiado do último estágio
COPY --from=builder /app/tangle-hornet-api .

# Definindo valores padrões para as variáveis de ambiente
ENV API_PORT=3000
ENV TANGLE_NODE_URL=127.0.0.1
ENV TANGLE_NODE_PORT=14265

# Cria que lê as variáveis de ambiente e escreve no arquivo de configuração da API
RUN echo '#!/bin/bash' > entrypoint.sh && \
    echo 'echo "apiPort = $API_PORT" > /etc/tangle-hornet.conf' >> entrypoint.sh && \
    echo 'echo "" >> /etc/tangle-hornet.conf' >> entrypoint.sh && \
    echo 'echo "nodeUrl = $TANGLE_NODE_URL" >> /etc/tangle-hornet.conf' >> entrypoint.sh && \
    echo 'echo "nodePort = $TANGLE_NODE_PORT" >> /etc/tangle-hornet.conf' >> entrypoint.sh && \
    echo 'exec ./tangle-hornet-api' >> entrypoint.sh

# Dá permissão de execução para o scrit criado
RUN chmod +x entrypoint.sh

# Garante que o diretório exista para receber o arquivo de configuração
RUN mkdir -p /etc/

EXPOSE 3000

# Define comando inicial de execução do sistema
ENTRYPOINT ["./entrypoint.sh"]

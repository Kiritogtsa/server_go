# Etapa de construção
FROM golang:1.22.4-alpine AS build

# Defina o diretório de trabalho para a aplicação
WORKDIR /app

# Copie os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do projeto usando go mod download
RUN go mod download

# Copie o código-fonte da aplicação para o diretório de trabalho
COPY . .

# Compile o binário da aplicação
RUN go build -o app

# Etapa final
FROM alpine:3.18

# Defina o diretório de trabalho para a aplicação
WORKDIR /app

# Copie o binário da etapa de construção
COPY --from=build /app/app /app/app

# Defina o comando de inicialização
CMD ["/app/app"]

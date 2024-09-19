# Usa una imagen base de Go para la construcción
FROM golang:1.23.0 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto de los archivos del proyecto
COPY . .

# Construye la aplicación
RUN go build -o myapp main.go

# Usa una imagen base más pequeña para la ejecución
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario construido desde la etapa anterior
COPY --from=builder /app/myapp .

# Expone el puerto en el que la aplicación escucha
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./myapp"]

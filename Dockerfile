FROM golang:1.22.5 as base

WORKDIR /app

COPY go.mod .

# In future if anyone added any dependency in go.mod file then it will download it
RUN go mod download

COPY . .

RUN go build -o main .

# Final stage with Distroless image

FROM gcr.io/distroless/base

COPY --from=base /app .

EXPOSE 90

CMD ["./main"]


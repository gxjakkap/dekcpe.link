FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/dekcpedotlink

FROM alpine:3.21 AS runner
RUN adduser -D -g '' fiber
COPY --from=builder /app/dekcpedotlink /dekcpedotlink
COPY --from=builder /app/views/* /views/
RUN chown fiber:fiber /dekcpedotlink
USER fiber
EXPOSE 3000
ENTRYPOINT ["/dekcpedotlink"]
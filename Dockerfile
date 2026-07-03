FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/spellbook .

FROM alpine:3
WORKDIR /app
COPY --from=build /out/spellbook ./spellbook
COPY 5e-SRD-Spells.json .
COPY static static
EXPOSE 1323
CMD ["./spellbook"]

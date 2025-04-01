# Build from source
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN make build/api

# Run tests
FROM build-stage AS run-test-stage
RUN make test 

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11:debug AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/bin/api /bin/api
COPY --from=build-stage /app/docs/openapi.yaml /docs/openapi.yaml

EXPOSE 4000 

USER nonroot:nonroot

ENTRYPOINT ["/bin/api"]

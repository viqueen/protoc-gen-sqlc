## protoc-gen-sqlc

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=viqueen_protoc-gen-sqlc&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=viqueen_protoc-gen-sqlc)

Protocol Buffers plugin to generate SQLC queries and schema from proto files.

---

### install it

- using **go**

```bash
go install github.com/<username>/protoc-gen-plugin/cmd@latest
```

---

### development setup

- install dependencies

```bash
go mod download
```

- codegen

```bash
./build.sh codegen
```

- build

```bash
./build.sh local
```
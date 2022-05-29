# Data Flusher

- Flushes all data in microservice env
  - It just **flushes** data in Postgresql
  - It **truncates** `product`, `variants` and `jobs_special_price_expiry` tables
  - It **doesn't remove** tables.
- Flushes Algolia
  - **Doesn't remove** indexes, replica indexs or other resources in **Algolia**
  - **Doesn't change** or **remove** index configurations

# Before Run

**Connect VPN** according to environment. Connect regular VPN for `dev` and `uat`. For prod connect `prod` VPN.

# Run

```go
ENV=dev|uat|prod go run main.go
```

or in **VS Code** just hit `F5`

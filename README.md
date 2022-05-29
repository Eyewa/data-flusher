# Data Flusher

- Flushes all data in microservice env
  - It just **flushes** all data in Postgresql (truncate table)
  - It **truncates** `product`, `variants` and `jobs_special_price_expiry` tables
  - It **doesn't remove** tables.
- Flushes Algolia
  - **Doesn't remove** indexes, replica indexs or other resources
  - **Doesn't change** or **remove** index configuration

# Before Run

**Connect VPN** according to environment. Connect regular VPN for `dev` and `uat`. For prod connect `prod` VPN.

# Run

Copy and paste any `.env.[dev|uat|prod]` file and keep it just `.env`. Then

```go
  go run main.go
```

or in **VS Code** just hit `F5`

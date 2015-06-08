# json-nullable
Serialize `NullString`, `NullFloat64`, `NullInt64`, `NullBool`, `NullTime` from SQL to JSON as it suppose

# Usage

It works just like `sql.NullString`, `sql.NullInt64`, but returns null value instead of `{ Valid: false, String: "" }`

```go
type Post struct {
    Title       string
    ParentID    nullable.NullString
    CreatedAt   nullable.NullTime
    DeletedAt   nullable.NullTime
}
```
If `DeletedAt` is null, it will be marshalized to JSON `null` instead of zero time: `0001-01-01 00:00:00 +0000 UTC`
If `ParentID` is null, it will be marshalized to JSON `null`

# Credits

I just took code from: https://github.com/soarcn/GoJsonSqlNullObj , add true `nil` serialization and `NullTime`.

# LICENSE

MIT

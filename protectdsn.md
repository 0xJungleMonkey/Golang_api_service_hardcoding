1. Store DSN securely in an environment variable
```
export DATABASE_DSN="root@tcp(127.0.0.1:3306)/timemanagementapp?charset=utf8mb4&parseTime=True&loc=Local"
```
2. retrieve the value of env:

```
echo $DATABASE_DSN
```
3. use it in golang code
```
os.Getenv("DATABASE_DSA")
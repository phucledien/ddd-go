# DDD in Go


## Structure

```bash
.
├── cmd
│   └── main.go
├── domain
│   ├── customer
│   │   ├── customer.go
│   │   ├── customer_test.go
│   │   ├── memory
│   │   │   ├── memory.go
│   │   │   └── memory_test.go
│   │   ├── mongo
│   │   │   └── mongo.go
│   │   └── repository.go
│   └── product
│       ├── memory
│       │   ├── memory.go
│       │   └── memory_test.go
│       ├── product.go
│       ├── product_test.go
│       └── repository.go
├── go.mod
├── go.sum
├── item.go
├── person.go
├── readme.md
├── services
│   ├── order
│   │   ├── order.go
│   │   └── order_test.go
│   └── tavern
│       ├── tavern.go
│       └── tavern_test.go
└── transaction.go

11 directories, 22 files
```

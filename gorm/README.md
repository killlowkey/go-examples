## Install Dependencies
```shell
go get gorm.io/gorm
# pure go sqlite dependencies
go get github.com/glebarez/sqlite
# cgo sqlite
# go get gorm.io/driver/sqlite
```

## users and book
```go
// User 用户实体
type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Age       string
	Address   string
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

// Book 书籍实体
type Book struct {
	ID        uint `gorm:"primarykey"`
	UserId    int
	Name      string
	Perice    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

```sql
CREATE TABLE `users` (
    `id` integer,
    `name` text,
    `age` text,
    `address` text,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
    PRIMARY KEY (`id`)
)
CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`)

CREATE TABLE `books` (
    `id` integer,
    `user_id` integer,
    `name` text,
    `perice` integer,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
     PRIMARY KEY (`id`),CONSTRAINT `fk_users_books` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
)
CREATE INDEX `idx_books_deleted_at` ON `books`(`deleted_at`)
```
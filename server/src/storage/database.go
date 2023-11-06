/*
 * @Date: 2023-10-25 07:38:19
 * @LastEditors: hxlh
 * @LastEditTime: 2023-10-28 09:19:04
 * @FilePath: /1024/server/src/storage/database.go
 */
package storage

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type BaseDB struct {
	DB *sql.DB
	RC *redis.Client
}

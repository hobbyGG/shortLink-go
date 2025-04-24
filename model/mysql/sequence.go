package mysqlse

import (
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 使用gorm实现取号功能

type Client struct {
	db *gorm.DB
}

type Sequence struct {
	ID        uint64    `gorm:"id"`
	Stub      string    `gorm:"stub"`
	Timestamp time.Time `gorm:"timestamp"`
}

func NewClient(datasource string) *Client {
	db, err := gorm.Open(mysql.Open(datasource))
	if err != nil {
		logx.Errorw("sequence.go", logx.Field("gorm.Open error", err))
		return nil
	}
	db.AutoMigrate(&Sequence{})
	return &Client{
		db: db,
	}
}

// Get 使用gorm进行取号
func (c Client) Get() (uint64, error) {
	tx := c.db.Session(&gorm.Session{
		PrepareStmt: true,
	})
	sqlStr := `replace into sequence(stub) values("a")`
	if err := tx.Model(&Sequence{}).Table("sequence").Exec(sqlStr).Error; err != nil {
		logx.Errorw("sequence.go error", logx.Field("tx.Exec", err))
		return 0, err
	}
	var id uint64
	if err := tx.Model(&Sequence{}).Table("sequence").Select("id").First(&id).Error; err != nil {
		logx.Errorw("sequence.go error", logx.Field("tx.Select", err))
		return 0, err
	}
	return id, nil
}

func InitBlackList(blConfig []string) map[string]struct{} {
	m := make(map[string]struct{}, len(blConfig))
	for _, v := range blConfig {
		m[v] = struct{}{}
	}
	return m
}

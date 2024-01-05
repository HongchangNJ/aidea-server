package data

import "github.com/mylxsw/eloquent/migrate"

func Migrate20240106DDL(m *migrate.Manager) {
	m.Schema("20240106-ddl").Table("chat_group_message", func(builder *migrate.Builder) {
		builder.Integer("first_letter_cost", false, true).Nullable(true).Comment("第一个字符响应耗时，单位微秒")
	})
}

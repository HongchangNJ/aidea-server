package data

import "github.com/mylxsw/eloquent/migrate"

func Migrate20240104DDL(m *migrate.Manager) {
	m.Schema("20240104-ddl").Table("chat_messages", func(builder *migrate.Builder) {
		builder.Integer("first_letter_cost", false, true).Nullable(true).Comment("第一个字符响应耗时，单位微秒")
		builder.Integer("total_cost", false, true).Nullable(true).Comment("总耗时，单位微秒")
	})

	m.Schema("20240104-ddl").Table("chat_group_message", func(builder *migrate.Builder) {
		builder.Integer("total_cost", false, true).Nullable(true).Comment("总耗时，单位微秒")
	})
}

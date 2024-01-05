package data

import "github.com/mylxsw/eloquent/migrate"

func Migrate20240105DDL(m *migrate.Manager) {
	m.Schema("20240105-ddl").Table("chat_messages", func(builder *migrate.Builder) {
		builder.TinyInteger("rating", false, true).Nullable(true).Comment("评价 1-5 （1-不好、5-好）")
	})

	m.Schema("20240105-ddl").Table("chat_group_message", func(builder *migrate.Builder) {
		builder.TinyInteger("rating", false, true).Nullable(true).Comment("评价 1-5 （1-不好、5-好）")
	})
}

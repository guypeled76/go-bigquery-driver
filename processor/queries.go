package processor

const (
	hasColumn     string = "__hasColumn($tableName,$columnName)"
	modifyColumn  string = "__modifyColumn($tableName,$columnName,$columnType)"
	hasTable      string = "__hasTable($tableName)"
	hasIndex      string = "__hasIndex($tableName,$indexName)"
	removeIndex   string = "__removeIndex($tableName,$indexName)"
	hasForeignKey string = "__hasForeignKey($tableName,$foreignKeyName)"
)

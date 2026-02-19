package awskinesisfirehose


// The sub-record type to deaggregate input records.
type SubRecordType string

const (
	// The records are JSON objects on a single line with no delimiter or newline-delimited (JSONL).
	SubRecordType_JSON SubRecordType = "JSON"
	// The records are delimited by a custom delimiter.
	SubRecordType_DELIMITED SubRecordType = "DELIMITED"
)


package awsfsx


// The permitted Lustre data compression algorithms.
//
// Example:
//   lustreConfiguration := map[string]lustreDataCompressionType{
//   	// ...
//   	"dataCompressionType": fsx.lustreDataCompressionType_LZ4,
//   }
//
type LustreDataCompressionType string

const (
	// `NONE` - (Default) Data compression is turned off when the file system is created.
	LustreDataCompressionType_NONE LustreDataCompressionType = "NONE"
	// `LZ4` - Data compression is turned on with the LZ4 algorithm.
	//
	// Note: When you turn data compression on for an existing file system, only newly written files are compressed. Existing files are not compressed.
	LustreDataCompressionType_LZ4 LustreDataCompressionType = "LZ4"
)


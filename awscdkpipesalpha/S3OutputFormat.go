package awscdkpipesalpha


// Log format for `S3LogDestination` logging configuration.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-s3logdestination.html#cfn-pipes-pipe-s3logdestination-outputformat
//
// Experimental.
type S3OutputFormat string

const (
	// Plain text.
	// Experimental.
	S3OutputFormat_PLAIN S3OutputFormat = "PLAIN"
	// JSON.
	// Experimental.
	S3OutputFormat_JSON S3OutputFormat = "JSON"
	// W3C extended log file format.
	// See: https://www.w3.org/TR/WD-logfile
	//
	// Experimental.
	S3OutputFormat_W3C S3OutputFormat = "W3C"
)


package awsecs


// Firelens configuration file type, s3 or file path.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig
// Experimental.
type FirelensConfigFileType string

const (
	// s3.
	// Experimental.
	FirelensConfigFileType_S3 FirelensConfigFileType = "S3"
	// fluentd.
	// Experimental.
	FirelensConfigFileType_FILE FirelensConfigFileType = "FILE"
)


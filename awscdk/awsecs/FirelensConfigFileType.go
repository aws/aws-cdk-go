package awsecs


// Firelens configuration file type, s3 or file path.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef-customconfig
type FirelensConfigFileType string

const (
	// s3.
	FirelensConfigFileType_S3 FirelensConfigFileType = "S3"
	// fluentd.
	FirelensConfigFileType_FILE FirelensConfigFileType = "FILE"
)


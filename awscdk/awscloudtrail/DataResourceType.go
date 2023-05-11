package awscloudtrail


// Resource type for a data event.
type DataResourceType string

const (
	// Data resource type for Lambda function.
	DataResourceType_LAMBDA_FUNCTION DataResourceType = "LAMBDA_FUNCTION"
	// Data resource type for S3 objects.
	DataResourceType_S3_OBJECT DataResourceType = "S3_OBJECT"
)


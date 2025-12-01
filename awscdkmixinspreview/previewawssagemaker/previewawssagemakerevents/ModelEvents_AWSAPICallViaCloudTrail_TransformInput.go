package previewawssagemakerevents


// Type definition for TransformInput.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformInput := &TransformInput{
//   	CompressionType: []*string{
//   		jsii.String("compressionType"),
//   	},
//   	ContentType: []*string{
//   		jsii.String("contentType"),
//   	},
//   	DataSource: &DataSource{
//   		S3DataSource: &S3DataSource{
//   			S3DataType: []*string{
//   				jsii.String("s3DataType"),
//   			},
//   			S3Uri: []*string{
//   				jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_TransformInput struct {
	// compressionType property.
	//
	// Specify an array of string values to match this event if the actual value of compressionType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompressionType *[]*string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// contentType property.
	//
	// Specify an array of string values to match this event if the actual value of contentType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContentType *[]*string `field:"optional" json:"contentType" yaml:"contentType"`
	// dataSource property.
	//
	// Specify an array of string values to match this event if the actual value of dataSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataSource *ModelEvents_AWSAPICallViaCloudTrail_DataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
}


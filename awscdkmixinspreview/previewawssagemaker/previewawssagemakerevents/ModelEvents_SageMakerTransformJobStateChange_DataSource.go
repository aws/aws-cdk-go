package previewawssagemakerevents


// Type definition for DataSource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSource := &DataSource{
//   	S3DataSource: &S3DataSource{
//   		S3DataType: []*string{
//   			jsii.String("s3DataType"),
//   		},
//   		S3Uri: []*string{
//   			jsii.String("s3Uri"),
//   		},
//   	},
//   }
//
// Experimental.
type ModelEvents_SageMakerTransformJobStateChange_DataSource struct {
	// S3DataSource property.
	//
	// Specify an array of string values to match this event if the actual value of S3DataSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3DataSource *ModelEvents_SageMakerTransformJobStateChange_S3DataSource `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}


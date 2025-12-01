package previewawssagemakerevents


// Type definition for S3DataSource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3DataSource := &S3DataSource{
//   	S3DataType: []*string{
//   		jsii.String("s3DataType"),
//   	},
//   	S3Uri: []*string{
//   		jsii.String("s3Uri"),
//   	},
//   }
//
// Experimental.
type ModelEvents_SageMakerTransformJobStateChange_S3DataSource struct {
	// S3DataType property.
	//
	// Specify an array of string values to match this event if the actual value of S3DataType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3DataType *[]*string `field:"optional" json:"s3DataType" yaml:"s3DataType"`
	// S3Uri property.
	//
	// Specify an array of string values to match this event if the actual value of S3Uri is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Uri *[]*string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}


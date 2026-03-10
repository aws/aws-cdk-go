package previewawscodepipelineevents


// Type definition for S3Location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3Location := &S3Location{
//   	Bucket: []*string{
//   		jsii.String("bucket"),
//   	},
//   	Key: []*string{
//   		jsii.String("key"),
//   	},
//   }
//
// Experimental.
type CodePipelineActionExecutionStateChange_S3Location struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *[]*string `field:"optional" json:"bucket" yaml:"bucket"`
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
}


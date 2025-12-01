package previewawssagemakerevents


// Type definition for OutputDataConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputDataConfig := &OutputDataConfig{
//   	RemoveJobNameFromS3OutputPath: []*string{
//   		jsii.String("removeJobNameFromS3OutputPath"),
//   	},
//   	S3OutputPath: []*string{
//   		jsii.String("s3OutputPath"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_OutputDataConfig struct {
	// removeJobNameFromS3OutputPath property.
	//
	// Specify an array of string values to match this event if the actual value of removeJobNameFromS3OutputPath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoveJobNameFromS3OutputPath *[]*string `field:"optional" json:"removeJobNameFromS3OutputPath" yaml:"removeJobNameFromS3OutputPath"`
	// s3OutputPath property.
	//
	// Specify an array of string values to match this event if the actual value of s3OutputPath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3OutputPath *[]*string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}


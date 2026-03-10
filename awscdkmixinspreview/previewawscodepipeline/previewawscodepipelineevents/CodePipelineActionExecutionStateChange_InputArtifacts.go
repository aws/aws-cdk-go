package previewawscodepipelineevents


// Type definition for InputArtifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputArtifacts := &InputArtifacts{
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	S3Location: &S3Location{
//   		Bucket: []*string{
//   			jsii.String("bucket"),
//   		},
//   		Key: []*string{
//   			jsii.String("key"),
//   		},
//   	},
//   }
//
// Experimental.
type CodePipelineActionExecutionStateChange_InputArtifacts struct {
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// s3location property.
	//
	// Specify an array of string values to match this event if the actual value of s3location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Location *CodePipelineActionExecutionStateChange_S3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}


package previewawssagemakerevents


// Type definition for TransformOutput.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformOutput := &TransformOutput{
//   	AssembleWith: []*string{
//   		jsii.String("assembleWith"),
//   	},
//   	S3OutputPath: []*string{
//   		jsii.String("s3OutputPath"),
//   	},
//   }
//
// Experimental.
type ModelEvents_SageMakerTransformJobStateChange_TransformOutput struct {
	// AssembleWith property.
	//
	// Specify an array of string values to match this event if the actual value of AssembleWith is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AssembleWith *[]*string `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// S3OutputPath property.
	//
	// Specify an array of string values to match this event if the actual value of S3OutputPath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3OutputPath *[]*string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}


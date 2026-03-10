package previewawssagemakerevents


// Type definition for ModelArtifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelArtifacts := &ModelArtifacts{
//   	S3ModelArtifacts: []*string{
//   		jsii.String("s3ModelArtifacts"),
//   	},
//   }
//
// Experimental.
type SageMakerTrainingJobStateChange_ModelArtifacts struct {
	// S3ModelArtifacts property.
	//
	// Specify an array of string values to match this event if the actual value of S3ModelArtifacts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3ModelArtifacts *[]*string `field:"optional" json:"s3ModelArtifacts" yaml:"s3ModelArtifacts"`
}


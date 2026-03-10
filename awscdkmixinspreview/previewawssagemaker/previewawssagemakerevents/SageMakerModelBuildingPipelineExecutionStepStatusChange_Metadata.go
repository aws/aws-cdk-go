package previewawssagemakerevents


// Type definition for Metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metadata := &Metadata{
//   	ProcessingJob: &ProcessingJob{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   	},
//   }
//
// Experimental.
type SageMakerModelBuildingPipelineExecutionStepStatusChange_Metadata struct {
	// processingJob property.
	//
	// Specify an array of string values to match this event if the actual value of processingJob is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProcessingJob *SageMakerModelBuildingPipelineExecutionStepStatusChange_ProcessingJob `field:"optional" json:"processingJob" yaml:"processingJob"`
}


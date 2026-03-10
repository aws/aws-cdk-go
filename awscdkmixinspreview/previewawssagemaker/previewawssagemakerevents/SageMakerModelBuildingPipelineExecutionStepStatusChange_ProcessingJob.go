package previewawssagemakerevents


// Type definition for ProcessingJob.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   processingJob := &ProcessingJob{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   }
//
// Experimental.
type SageMakerModelBuildingPipelineExecutionStepStatusChange_ProcessingJob struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
}


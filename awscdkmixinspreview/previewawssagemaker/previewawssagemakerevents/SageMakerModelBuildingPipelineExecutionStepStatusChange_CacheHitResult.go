package previewawssagemakerevents


// Type definition for CacheHitResult.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheHitResult := &CacheHitResult{
//   	SourcePipelineExecutionArn: []*string{
//   		jsii.String("sourcePipelineExecutionArn"),
//   	},
//   }
//
// Experimental.
type SageMakerModelBuildingPipelineExecutionStepStatusChange_CacheHitResult struct {
	// sourcePipelineExecutionArn property.
	//
	// Specify an array of string values to match this event if the actual value of sourcePipelineExecutionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourcePipelineExecutionArn *[]*string `field:"optional" json:"sourcePipelineExecutionArn" yaml:"sourcePipelineExecutionArn"`
}


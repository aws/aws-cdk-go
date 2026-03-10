package previewawscodepipelineevents


// Type definition for ExecutionResult.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   executionResult := &ExecutionResult{
//   	ExternalExecutionId: []*string{
//   		jsii.String("externalExecutionId"),
//   	},
//   	ExternalExecutionSummary: []*string{
//   		jsii.String("externalExecutionSummary"),
//   	},
//   	ExternalExecutionUrl: []*string{
//   		jsii.String("externalExecutionUrl"),
//   	},
//   }
//
// Experimental.
type CodePipelineActionExecutionStateChange_ExecutionResult struct {
	// external-execution-id property.
	//
	// Specify an array of string values to match this event if the actual value of external-execution-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExternalExecutionId *[]*string `field:"optional" json:"externalExecutionId" yaml:"externalExecutionId"`
	// external-execution-summary property.
	//
	// Specify an array of string values to match this event if the actual value of external-execution-summary is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExternalExecutionSummary *[]*string `field:"optional" json:"externalExecutionSummary" yaml:"externalExecutionSummary"`
	// external-execution-url property.
	//
	// Specify an array of string values to match this event if the actual value of external-execution-url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExternalExecutionUrl *[]*string `field:"optional" json:"externalExecutionUrl" yaml:"externalExecutionUrl"`
}


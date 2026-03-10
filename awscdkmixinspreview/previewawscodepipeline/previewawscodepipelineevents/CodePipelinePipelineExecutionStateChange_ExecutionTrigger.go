package previewawscodepipelineevents


// Type definition for execution-trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   executionTrigger := &ExecutionTrigger{
//   	TriggerDetail: []*string{
//   		jsii.String("triggerDetail"),
//   	},
//   	TriggerType: []*string{
//   		jsii.String("triggerType"),
//   	},
//   }
//
// Experimental.
type CodePipelinePipelineExecutionStateChange_ExecutionTrigger struct {
	// trigger-detail property.
	//
	// Specify an array of string values to match this event if the actual value of trigger-detail is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TriggerDetail *[]*string `field:"optional" json:"triggerDetail" yaml:"triggerDetail"`
	// trigger-type property.
	//
	// Specify an array of string values to match this event if the actual value of trigger-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TriggerType *[]*string `field:"optional" json:"triggerType" yaml:"triggerType"`
}


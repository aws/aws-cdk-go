package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBatchingConditionProperty := &eventBatchingConditionProperty{
//   	batchSize: jsii.Number(123),
//
//   	// the properties below are optional
//   	batchWindow: jsii.Number(123),
//   }
//
type CfnTrigger_EventBatchingConditionProperty struct {
	// `CfnTrigger.EventBatchingConditionProperty.BatchSize`.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// `CfnTrigger.EventBatchingConditionProperty.BatchWindow`.
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}


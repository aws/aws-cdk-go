package awsglue


// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
//
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
	// Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// Window of time in seconds after which EventBridge event trigger fires.
	//
	// Window starts when first event is received.
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}


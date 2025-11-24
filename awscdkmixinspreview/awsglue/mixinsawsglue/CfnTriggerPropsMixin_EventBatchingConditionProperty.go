package mixinsawsglue


// Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventBatchingConditionProperty := &EventBatchingConditionProperty{
//   	BatchSize: jsii.Number(123),
//   	BatchWindow: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-eventbatchingcondition.html
//
type CfnTriggerPropsMixin_EventBatchingConditionProperty struct {
	// Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-eventbatchingcondition.html#cfn-glue-trigger-eventbatchingcondition-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Window of time in seconds after which EventBridge event trigger fires.
	//
	// Window starts when first event is received.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-eventbatchingcondition.html#cfn-glue-trigger-eventbatchingcondition-batchwindow
	//
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}


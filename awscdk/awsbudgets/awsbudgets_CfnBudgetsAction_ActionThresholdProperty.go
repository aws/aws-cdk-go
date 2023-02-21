package awsbudgets


// The trigger threshold of the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionThresholdProperty := &actionThresholdProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
//
type CfnBudgetsAction_ActionThresholdProperty struct {
	// The type of threshold for a notification.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The threshold of a notification.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementConstraintProperty := &placementConstraintProperty{
//   	expression: jsii.String("expression"),
//   	type: jsii.String("type"),
//   }
//
type CfnSchedule_PlacementConstraintProperty struct {
	// `CfnSchedule.PlacementConstraintProperty.Expression`.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// `CfnSchedule.PlacementConstraintProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


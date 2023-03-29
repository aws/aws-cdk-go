package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardProperty := &ForwardProperty{
//   	TargetGroups: []interface{}{
//   		&WeightedTargetGroupProperty{
//   			TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   			// the properties below are optional
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnListener_ForwardProperty struct {
	// `CfnListener.ForwardProperty.TargetGroups`.
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}


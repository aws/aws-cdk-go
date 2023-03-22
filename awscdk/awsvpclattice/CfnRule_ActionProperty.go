package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	Forward: &ForwardProperty{
//   		TargetGroups: []interface{}{
//   			&WeightedTargetGroupProperty{
//   				TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   				// the properties below are optional
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnRule_ActionProperty struct {
	// `CfnRule.ActionProperty.Forward`.
	Forward interface{} `field:"required" json:"forward" yaml:"forward"`
}


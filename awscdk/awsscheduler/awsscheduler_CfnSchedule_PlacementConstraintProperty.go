package awsscheduler


// An object representing a constraint on task placement.
//
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
	// A cluster query language expression to apply to the constraint.
	//
	// You cannot specify an expression if the constraint type is `distinctInstance` . For more information, see [Cluster query language](https://docs.aws.amazon.com/latest/developerguide/cluster-query-language.html) in the *Amazon ECS Developer Guide* .
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The type of constraint.
	//
	// Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


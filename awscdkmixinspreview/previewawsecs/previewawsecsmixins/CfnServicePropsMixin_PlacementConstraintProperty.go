package previewawsecsmixins


// An object representing a constraint on task placement.
//
// For more information, see [Task placement constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > If you're using the Fargate launch type, task placement constraints aren't supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   placementConstraintProperty := &PlacementConstraintProperty{
//   	Expression: jsii.String("expression"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html
//
type CfnServicePropsMixin_PlacementConstraintProperty struct {
	// A cluster query language expression to apply to the constraint.
	//
	// The expression can have a maximum length of 2000 characters. You can't specify an expression if the constraint type is `distinctInstance` . For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The type of constraint.
	//
	// Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


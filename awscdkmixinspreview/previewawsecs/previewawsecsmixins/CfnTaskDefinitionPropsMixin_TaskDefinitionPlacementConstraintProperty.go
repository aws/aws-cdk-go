package previewawsecsmixins


// The constraint on task placement in the task definition.
//
// For more information, see [Task placement constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > Task placement constraints aren't supported for tasks run on AWS Fargate .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskDefinitionPlacementConstraintProperty := &TaskDefinitionPlacementConstraintProperty{
//   	Expression: jsii.String("expression"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html
//
type CfnTaskDefinitionPropsMixin_TaskDefinitionPlacementConstraintProperty struct {
	// A cluster query language expression to apply to the constraint.
	//
	// For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The type of constraint.
	//
	// The `MemberOf` constraint restricts selection to be from a group of valid candidates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


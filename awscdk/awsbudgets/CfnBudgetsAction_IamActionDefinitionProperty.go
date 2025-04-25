package awsbudgets


// The AWS Identity and Access Management ( IAM ) action definition details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamActionDefinitionProperty := &IamActionDefinitionProperty{
//   	PolicyArn: jsii.String("policyArn"),
//
//   	// the properties below are optional
//   	Groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	Roles: []*string{
//   		jsii.String("roles"),
//   	},
//   	Users: []*string{
//   		jsii.String("users"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-iamactiondefinition.html
//
type CfnBudgetsAction_IamActionDefinitionProperty struct {
	// The Amazon Resource Name (ARN) of the policy to be attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-iamactiondefinition.html#cfn-budgets-budgetsaction-iamactiondefinition-policyarn
	//
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// A list of groups to be attached.
	//
	// There must be at least one group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-iamactiondefinition.html#cfn-budgets-budgetsaction-iamactiondefinition-groups
	//
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// A list of roles to be attached.
	//
	// There must be at least one role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-iamactiondefinition.html#cfn-budgets-budgetsaction-iamactiondefinition-roles
	//
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// A list of users to be attached.
	//
	// There must be at least one user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-iamactiondefinition.html#cfn-budgets-budgetsaction-iamactiondefinition-users
	//
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}


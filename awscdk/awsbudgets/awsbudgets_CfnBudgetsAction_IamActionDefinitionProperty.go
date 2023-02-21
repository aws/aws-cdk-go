package awsbudgets


// The AWS Identity and Access Management ( IAM ) action definition details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamActionDefinitionProperty := &iamActionDefinitionProperty{
//   	policyArn: jsii.String("policyArn"),
//
//   	// the properties below are optional
//   	groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	roles: []*string{
//   		jsii.String("roles"),
//   	},
//   	users: []*string{
//   		jsii.String("users"),
//   	},
//   }
//
type CfnBudgetsAction_IamActionDefinitionProperty struct {
	// The Amazon Resource Name (ARN) of the policy to be attached.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// A list of groups to be attached.
	//
	// There must be at least one group.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// A list of roles to be attached.
	//
	// There must be at least one role.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// A list of users to be attached.
	//
	// There must be at least one user.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}


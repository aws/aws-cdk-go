package awsiam


// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnPolicyProps := &cfnPolicyProps{
//   	policyDocument: policyDocument,
//   	policyName: jsii.String("policyName"),
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
type CfnPolicyProps struct {
	// The policy document.
	//
	// You must provide policies in JSON format in IAM. However, for AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
	//
	// - Any printable ASCII character ranging from the space character ( `\ u0020` ) through the end of the ASCII character range
	// - The printable characters in the Basic Latin and Latin-1 Supplement character set (through `\ u00FF` )
	// - The special characters tab ( `\ u0009` ), line feed ( `\ u000A` ), and carriage return ( `\ u000D` ).
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the policy document.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the group to associate the policy with.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// The name of the role to associate the policy with.
	//
	// This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// > If an external policy (such as `AWS::IAM::Policy` or `AWS::IAM::ManagedPolicy` ) has a `Ref` to a role and if a resource (such as `AWS::ECS::Service` ) also has a `Ref` to the same role, add a `DependsOn` attribute to the resource to make the resource depend on the external policy. This dependency ensures that the role's policy is available throughout the resource's lifecycle. For example, when you delete a stack with an `AWS::ECS::Service` resource, the `DependsOn` attribute ensures that AWS CloudFormation deletes the `AWS::ECS::Service` resource before deleting its role's policy.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// The name of the user to associate the policy with.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}


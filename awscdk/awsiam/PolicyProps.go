package awsiam


// Properties for defining an IAM inline policy document.
//
// Example:
//   var books resource
//   var iamUser user
//
//
//   getBooks := books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	AuthorizationType: apigateway.AuthorizationType_IAM,
//   })
//
//   iamUser.AttachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowBooks"), &PolicyProps{
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("execute-api:Invoke"),
//   			},
//   			Effect: iam.Effect_ALLOW,
//   			Resources: []*string{
//   				getBooks.methodArn,
//   			},
//   		}),
//   	},
//   }))
//
type PolicyProps struct {
	// Initial PolicyDocument to use for this Policy.
	//
	// If omited, any
	// `PolicyStatement` provided in the `statements` property will be applied
	// against the empty default `PolicyDocument`.
	// Default: - An empty policy.
	//
	Document PolicyDocument `field:"optional" json:"document" yaml:"document"`
	// Force creation of an `AWS::IAM::Policy`.
	//
	// Unless set to `true`, this `Policy` construct will not materialize to an
	// `AWS::IAM::Policy` CloudFormation resource in case it would have no effect
	// (for example, if it remains unattached to an IAM identity or if it has no
	// statements). This is generally desired behavior, since it prevents
	// creating invalid--and hence undeployable--CloudFormation templates.
	//
	// In cases where you know the policy must be created and it is actually
	// an error if no statements have been added to it, you can set this to `true`.
	// Default: false.
	//
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// Groups to attach this policy to.
	//
	// You can also use `attachToGroup(group)` to attach this policy to a group.
	// Default: - No groups.
	//
	Groups *[]IGroup `field:"optional" json:"groups" yaml:"groups"`
	// The name of the policy.
	//
	// If you specify multiple policies for an entity,
	// specify unique names. For example, if you specify a list of policies for
	// an IAM role, each policy must have a unique name.
	// Default: - Uses the logical ID of the policy resource, which is ensured
	// to be unique within the stack.
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Roles to attach this policy to.
	//
	// You can also use `attachToRole(role)` to attach this policy to a role.
	// Default: - No roles.
	//
	Roles *[]IRole `field:"optional" json:"roles" yaml:"roles"`
	// Initial set of permissions to add to this policy document.
	//
	// You can also use `addStatements(...statement)` to add permissions later.
	// Default: - No statements.
	//
	Statements *[]PolicyStatement `field:"optional" json:"statements" yaml:"statements"`
	// Users to attach this policy to.
	//
	// You can also use `attachToUser(user)` to attach this policy to a user.
	// Default: - No users.
	//
	Users *[]IUser `field:"optional" json:"users" yaml:"users"`
}


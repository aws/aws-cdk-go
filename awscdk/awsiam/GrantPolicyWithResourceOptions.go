package awsiam


// Options for a grant operation that directly adds a policy statement to a resource.
//
// This differs from GrantWithResourceOptions in that it requires a pre-constructed
// PolicyStatement rather than constructing one from individual permissions.
// Use this when you need fine-grained control over the initial policy statement's contents.
//
// Example:
//   var grantee iGrantable
//   var actions []*string
//   var resourceArns []*string
//   var bucket bucket
//
//
//   statement := iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Actions: actions,
//   	Principals: []iPrincipal{
//   		iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   	},
//   	Conditions: map[string]interface{}{
//   		"StringEquals": map[string]*string{
//   			"aws:SourceAccount": awscdk.*stack_of(this).account,
//   		},
//   	},
//   })
//   iam.Grant_AddStatementToResourcePolicy(&GrantPolicyWithResourceOptions{
//   	Grantee: grantee,
//   	Actions: actions,
//   	ResourceArns: resourceArns,
//   	Resource: bucket,
//   	Statement: statement,
//   })
//
type GrantPolicyWithResourceOptions struct {
	// The actions to grant.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The principal to grant to.
	// Default: if principal is undefined, no work is done.
	//
	Grantee IGrantable `field:"required" json:"grantee" yaml:"grantee"`
	// The resource ARNs to grant to.
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// Any conditions to attach to the grant.
	// Default: - No conditions.
	//
	Conditions *map[string]*map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The resource with a resource policy.
	//
	// The statement will be added to the resource policy if it couldn't be
	// added to the principal policy.
	Resource IResourceWithPolicy `field:"required" json:"resource" yaml:"resource"`
	// When referring to the resource in a resource policy, use this as ARN.
	//
	// (Depending on the resource type, this needs to be '*' in a resource policy).
	// Default: Same as regular resource ARNs.
	//
	ResourceSelfArns *[]*string `field:"optional" json:"resourceSelfArns" yaml:"resourceSelfArns"`
	// The policy statement to add to the resource's policy.
	//
	// This statement will be passed to the resource's addToResourcePolicy method.
	// The actual handling of the statement depends on the specific IResourceWithPolicy
	// implementation.
	Statement PolicyStatement `field:"required" json:"statement" yaml:"statement"`
}


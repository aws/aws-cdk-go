package awsiam


// Options for a grant operation to both identity and resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//   var grantable iGrantable
//   var principal iPrincipal
//   var resourceWithPolicy iResourceWithPolicy
//
//   grantOnPrincipalAndResourceOptions := &GrantOnPrincipalAndResourceOptions{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	Grantee: grantable,
//   	Resource: resourceWithPolicy,
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//
//   	// the properties below are optional
//   	Conditions: map[string]map[string]interface{}{
//   		"conditionsKey": map[string]interface{}{
//   			"conditionsKey": conditions,
//   		},
//   	},
//   	ResourcePolicyPrincipal: principal,
//   	ResourceSelfArns: []*string{
//   		jsii.String("resourceSelfArns"),
//   	},
//   }
//
type GrantOnPrincipalAndResourceOptions struct {
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
	// The statement will always be added to the resource policy.
	Resource IResourceWithPolicy `field:"required" json:"resource" yaml:"resource"`
	// The principal to use in the statement for the resource policy.
	// Default: - the principal of the grantee will be used.
	//
	ResourcePolicyPrincipal IPrincipal `field:"optional" json:"resourcePolicyPrincipal" yaml:"resourcePolicyPrincipal"`
	// When referring to the resource in a resource policy, use this as ARN.
	//
	// (Depending on the resource type, this needs to be '*' in a resource policy).
	// Default: Same as regular resource ARNs.
	//
	ResourceSelfArns *[]*string `field:"optional" json:"resourceSelfArns" yaml:"resourceSelfArns"`
}


package awsiam


// Options for a grant operation to both identity and resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var grantable iGrantable
//   var principal iPrincipal
//   var resourceWithPolicy iResourceWithPolicy
//
//   grantOnPrincipalAndResourceOptions := &grantOnPrincipalAndResourceOptions{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	grantee: grantable,
//   	resource: resourceWithPolicy,
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//
//   	// the properties below are optional
//   	resourcePolicyPrincipal: principal,
//   	resourceSelfArns: []*string{
//   		jsii.String("resourceSelfArns"),
//   	},
//   }
//
// Experimental.
type GrantOnPrincipalAndResourceOptions struct {
	// The actions to grant.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The principal to grant to.
	// Experimental.
	Grantee IGrantable `field:"required" json:"grantee" yaml:"grantee"`
	// The resource ARNs to grant to.
	// Experimental.
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// The resource with a resource policy.
	//
	// The statement will always be added to the resource policy.
	// Experimental.
	Resource IResourceWithPolicy `field:"required" json:"resource" yaml:"resource"`
	// The principal to use in the statement for the resource policy.
	// Experimental.
	ResourcePolicyPrincipal IPrincipal `field:"optional" json:"resourcePolicyPrincipal" yaml:"resourcePolicyPrincipal"`
	// When referring to the resource in a resource policy, use this as ARN.
	//
	// (Depending on the resource type, this needs to be '*' in a resource policy).
	// Experimental.
	ResourceSelfArns *[]*string `field:"optional" json:"resourceSelfArns" yaml:"resourceSelfArns"`
}


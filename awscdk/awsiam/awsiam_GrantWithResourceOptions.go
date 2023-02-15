package awsiam


// Options for a grant operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//   var grantable iGrantable
//   var resourceWithPolicy iResourceWithPolicy
//
//   grantWithResourceOptions := &grantWithResourceOptions{
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
//   	conditions: map[string]map[string]interface{}{
//   		"conditionsKey": map[string]interface{}{
//   			"conditionsKey": conditions,
//   		},
//   	},
//   	resourceSelfArns: []*string{
//   		jsii.String("resourceSelfArns"),
//   	},
//   }
//
type GrantWithResourceOptions struct {
	// The actions to grant.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The principal to grant to.
	Grantee IGrantable `field:"required" json:"grantee" yaml:"grantee"`
	// The resource ARNs to grant to.
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// Any conditions to attach to the grant.
	Conditions *map[string]*map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The resource with a resource policy.
	//
	// The statement will be added to the resource policy if it couldn't be
	// added to the principal policy.
	Resource IResourceWithPolicy `field:"required" json:"resource" yaml:"resource"`
	// When referring to the resource in a resource policy, use this as ARN.
	//
	// (Depending on the resource type, this needs to be '*' in a resource policy).
	ResourceSelfArns *[]*string `field:"optional" json:"resourceSelfArns" yaml:"resourceSelfArns"`
}


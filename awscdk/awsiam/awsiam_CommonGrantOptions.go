package awsiam


// Basic options for a grant operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//   var grantable iGrantable
//
//   commonGrantOptions := &commonGrantOptions{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	grantee: grantable,
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
//   }
//
type CommonGrantOptions struct {
	// The actions to grant.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The principal to grant to.
	Grantee IGrantable `field:"required" json:"grantee" yaml:"grantee"`
	// The resource ARNs to grant to.
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// Any conditions to attach to the grant.
	Conditions *map[string]*map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}


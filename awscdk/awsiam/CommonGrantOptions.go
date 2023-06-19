package awsiam


// Basic options for a grant operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var grantable iGrantable
//
//   commonGrantOptions := &CommonGrantOptions{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	Grantee: grantable,
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   }
//
// Experimental.
type CommonGrantOptions struct {
	// The actions to grant.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The principal to grant to.
	// Experimental.
	Grantee IGrantable `field:"required" json:"grantee" yaml:"grantee"`
	// The resource ARNs to grant to.
	// Experimental.
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
}


package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for a grant operation that only applies to principals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//   var grantable iGrantable
//
//   grantOnPrincipalOptions := &grantOnPrincipalOptions{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	grantee: grantable,
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//
//   	// the properties below are optional
//   	scope: construct,
//   }
//
// Experimental.
type GrantOnPrincipalOptions struct {
	// The actions to grant.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The principal to grant to.
	// Experimental.
	Grantee IGrantable `field:"required" json:"grantee" yaml:"grantee"`
	// The resource ARNs to grant to.
	// Experimental.
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// Construct to report warnings on in case grant could not be registered.
	// Experimental.
	Scope awscdk.IConstruct `field:"optional" json:"scope" yaml:"scope"`
}


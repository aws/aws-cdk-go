// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Options to the resolve() operation.
//
// NOT the same as the ResolveContext; ResolveContext is exposed to Token
// implementors and resolution hooks, whereas this struct is just to bundle
// a number of things that would otherwise be arguments to resolve() in a
// readable way.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var tokenResolver iTokenResolver
//
//   resolveOptions := &resolveOptions{
//   	resolver: tokenResolver,
//   	scope: construct,
//
//   	// the properties below are optional
//   	preparing: jsii.Boolean(false),
//   	removeEmpty: jsii.Boolean(false),
//   }
//
type ResolveOptions struct {
	// The resolver to apply to any resolvable tokens found.
	Resolver ITokenResolver `field:"required" json:"resolver" yaml:"resolver"`
	// The scope from which resolution is performed.
	Scope constructs.IConstruct `field:"required" json:"scope" yaml:"scope"`
	// Whether the resolution is being executed during the prepare phase or not.
	Preparing *bool `field:"optional" json:"preparing" yaml:"preparing"`
	// Whether to remove undefined elements from arrays and objects when resolving.
	RemoveEmpty *bool `field:"optional" json:"removeEmpty" yaml:"removeEmpty"`
}


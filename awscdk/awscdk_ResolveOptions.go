// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"github.com/aws/constructs-go/constructs/v3"
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
//   import constructs "github.com/aws/constructs-go/constructs"
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//   var tokenResolver iTokenResolver
//
//   resolveOptions := &ResolveOptions{
//   	Resolver: tokenResolver,
//   	Scope: construct,
//
//   	// the properties below are optional
//   	Preparing: jsii.Boolean(false),
//   	RemoveEmpty: jsii.Boolean(false),
//   }
//
// Experimental.
type ResolveOptions struct {
	// The resolver to apply to any resolvable tokens found.
	// Experimental.
	Resolver ITokenResolver `field:"required" json:"resolver" yaml:"resolver"`
	// The scope from which resolution is performed.
	// Experimental.
	Scope constructs.IConstruct `field:"required" json:"scope" yaml:"scope"`
	// Whether the resolution is being executed during the prepare phase or not.
	// Experimental.
	Preparing *bool `field:"optional" json:"preparing" yaml:"preparing"`
	// Whether to remove undefined elements from arrays and objects when resolving.
	// Experimental.
	RemoveEmpty *bool `field:"optional" json:"removeEmpty" yaml:"removeEmpty"`
}


package awsiam

import (
	"github.com/aws/constructs-go/constructs/v3"
)

// Properties for an UnknownPrincipal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import constructs "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//
//   unknownPrincipalProps := &UnknownPrincipalProps{
//   	Resource: construct,
//   }
//
// Experimental.
type UnknownPrincipalProps struct {
	// The resource the role proxy is for.
	// Experimental.
	Resource constructs.IConstruct `field:"required" json:"resource" yaml:"resource"`
}


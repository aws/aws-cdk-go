package awsiam

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for an UnknownPrincipal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//
//   unknownPrincipalProps := &unknownPrincipalProps{
//   	resource: construct,
//   }
//
type UnknownPrincipalProps struct {
	// The resource the role proxy is for.
	Resource constructs.IConstruct `field:"required" json:"resource" yaml:"resource"`
}


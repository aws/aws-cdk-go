package awscdk

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// This defines the values needed for Injection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct Construct
//
//   injectionContext := &InjectionContext{
//   	Id: jsii.String("id"),
//   	Scope: construct,
//   }
//
type InjectionContext struct {
	// id from the Construct constructor.
	Id *string `field:"required" json:"id" yaml:"id"`
	// scope from the  constructor.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}


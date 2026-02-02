package core

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a successful mixin application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct Construct
//   var mixin Mixin
//
//   mixinApplication := &MixinApplication{
//   	Construct: construct,
//   	Mixin: mixin,
//   }
//
// Experimental.
type MixinApplication struct {
	// The construct the mixin was applied to.
	// Experimental.
	Construct constructs.IConstruct `field:"required" json:"construct" yaml:"construct"`
	// The mixin that was applied.
	// Experimental.
	Mixin IMixin `field:"required" json:"mixin" yaml:"mixin"`
}


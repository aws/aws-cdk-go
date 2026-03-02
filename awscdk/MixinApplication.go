package awscdk

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a successful mixin application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct Construct
//   var mixin IMixin
//
//   mixinApplication := &MixinApplication{
//   	Construct: construct,
//   	Mixin: mixin,
//   }
//
type MixinApplication struct {
	// The construct the mixin was applied to.
	Construct constructs.IConstruct `field:"required" json:"construct" yaml:"construct"`
	// The mixin that was applied.
	Mixin constructs.IMixin `field:"required" json:"mixin" yaml:"mixin"`
}


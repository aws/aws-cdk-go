package triggers

import (
	"github.com/aws/constructs-go/constructs/v3"
)

// Options for `Trigger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import constructs "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//
//   triggerOptions := &TriggerOptions{
//   	ExecuteAfter: []*construct{
//   		construct,
//   	},
//   	ExecuteBefore: []*construct{
//   		construct,
//   	},
//   	ExecuteOnHandlerChange: jsii.Boolean(false),
//   }
//
// Experimental.
type TriggerOptions struct {
	// Adds trigger dependencies. Execute this trigger only after these construct scopes have been provisioned.
	//
	// You can also use `trigger.executeAfter()` to add additional dependencies.
	// Experimental.
	ExecuteAfter *[]constructs.Construct `field:"optional" json:"executeAfter" yaml:"executeAfter"`
	// Adds this trigger as a dependency on other constructs.
	//
	// This means that this
	// trigger will get executed *before* the given construct(s).
	//
	// You can also use `trigger.executeBefore()` to add additional dependants.
	// Experimental.
	ExecuteBefore *[]constructs.Construct `field:"optional" json:"executeBefore" yaml:"executeBefore"`
	// Re-executes the trigger every time the handler changes.
	//
	// This implies that the trigger is associated with the `currentVersion` of
	// the handler, which gets recreated every time the handler or its
	// configuration is updated.
	// Experimental.
	ExecuteOnHandlerChange *bool `field:"optional" json:"executeOnHandlerChange" yaml:"executeOnHandlerChange"`
}


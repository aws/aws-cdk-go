package interfacesawsglue


// A reference to a Trigger resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerReference := &TriggerReference{
//   	TriggerName: jsii.String("triggerName"),
//   }
//
type TriggerReference struct {
	// The Name of the Trigger resource.
	TriggerName *string `field:"required" json:"triggerName" yaml:"triggerName"`
}


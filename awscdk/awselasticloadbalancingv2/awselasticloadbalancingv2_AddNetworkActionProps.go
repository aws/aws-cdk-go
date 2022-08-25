package awselasticloadbalancingv2


// Properties for adding a new action to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkListenerAction networkListenerAction
//
//   addNetworkActionProps := &addNetworkActionProps{
//   	action: networkListenerAction,
//   }
//
type AddNetworkActionProps struct {
	// Action to perform.
	Action NetworkListenerAction `field:"required" json:"action" yaml:"action"`
}


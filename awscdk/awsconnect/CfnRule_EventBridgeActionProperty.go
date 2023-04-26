package awsconnect


// The EventBridge action definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeActionProperty := &EventBridgeActionProperty{
//   	Name: jsii.String("name"),
//   }
//
type CfnRule_EventBridgeActionProperty struct {
	// The name.
	Name *string `field:"required" json:"name" yaml:"name"`
}


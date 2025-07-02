package awsappmesh


// Properties for a VirtualRouter listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterListenerConfig := &VirtualRouterListenerConfig{
//   	Listener: &VirtualRouterListenerProperty{
//   		PortMapping: &PortMappingProperty{
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//   }
//
type VirtualRouterListenerConfig struct {
	// Single listener config for a VirtualRouter.
	Listener *CfnVirtualRouter_VirtualRouterListenerProperty `field:"required" json:"listener" yaml:"listener"`
}


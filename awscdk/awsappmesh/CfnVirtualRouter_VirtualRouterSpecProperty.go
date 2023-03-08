package awsappmesh


// An object that represents the specification of a virtual router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterSpecProperty := &VirtualRouterSpecProperty{
//   	Listeners: []interface{}{
//   		&VirtualRouterListenerProperty{
//   			PortMapping: &PortMappingProperty{
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualRouter_VirtualRouterSpecProperty struct {
	// The listeners that the virtual router is expected to receive inbound traffic from.
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
}


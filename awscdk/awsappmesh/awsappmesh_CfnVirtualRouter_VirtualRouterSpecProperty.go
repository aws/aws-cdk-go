package awsappmesh


// An object that represents the specification of a virtual router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterSpecProperty := &virtualRouterSpecProperty{
//   	listeners: []interface{}{
//   		&virtualRouterListenerProperty{
//   			portMapping: &portMappingProperty{
//   				port: jsii.Number(123),
//   				protocol: jsii.String("protocol"),
//   			},
//   		},
//   	},
//   }
//
type CfnVirtualRouter_VirtualRouterSpecProperty struct {
	// The listeners that the virtual router is expected to receive inbound traffic from.
	//
	// You can specify one listener.
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
}


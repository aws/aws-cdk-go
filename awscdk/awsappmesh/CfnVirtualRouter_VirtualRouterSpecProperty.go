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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-virtualrouterspec.html
//
type CfnVirtualRouter_VirtualRouterSpecProperty struct {
	// The listeners that the virtual router is expected to receive inbound traffic from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-virtualrouterspec.html#cfn-appmesh-virtualrouter-virtualrouterspec-listeners
	//
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
}


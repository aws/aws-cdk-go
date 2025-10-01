package awsappmesh


// An object that represents a virtual router listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterListenerProperty := &VirtualRouterListenerProperty{
//   	PortMapping: &PortMappingProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-virtualrouterlistener.html
//
type CfnVirtualRouter_VirtualRouterListenerProperty struct {
	// The port mapping information for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-virtualrouterlistener.html#cfn-appmesh-virtualrouter-virtualrouterlistener-portmapping
	//
	PortMapping interface{} `field:"required" json:"portMapping" yaml:"portMapping"`
}


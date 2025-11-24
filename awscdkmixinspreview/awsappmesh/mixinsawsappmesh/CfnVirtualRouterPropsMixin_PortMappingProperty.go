package mixinsawsappmesh


// An object representing a virtual router listener port mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portMappingProperty := &PortMappingProperty{
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-portmapping.html
//
type CfnVirtualRouterPropsMixin_PortMappingProperty struct {
	// The port used for the port mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-portmapping.html#cfn-appmesh-virtualrouter-portmapping-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol used for the port mapping.
	//
	// Specify one protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-portmapping.html#cfn-appmesh-virtualrouter-portmapping-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


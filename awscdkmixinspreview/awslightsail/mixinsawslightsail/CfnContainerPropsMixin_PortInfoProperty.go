package mixinsawslightsail


// `PortInfo` is a property of the [Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html) property. It describes the ports to open and the protocols to use for a container on a Amazon Lightsail container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portInfoProperty := &PortInfoProperty{
//   	Port: jsii.String("port"),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-portinfo.html
//
type CfnContainerPropsMixin_PortInfoProperty struct {
	// The open firewall ports of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-portinfo.html#cfn-lightsail-container-portinfo-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol name for the open ports.
	//
	// *Allowed values* : `HTTP` | `HTTPS` | `TCP` | `UDP`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-portinfo.html#cfn-lightsail-container-portinfo-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


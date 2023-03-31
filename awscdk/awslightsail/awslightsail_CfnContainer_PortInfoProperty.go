package awslightsail


// `PortInfo` is a property of the [Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html) property. It describes the ports to open and the protocols to use for a container on a Amazon Lightsail container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portInfoProperty := &portInfoProperty{
//   	port: jsii.String("port"),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnContainer_PortInfoProperty struct {
	// The open firewall ports of the container.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol name for the open ports.
	//
	// *Allowed values* : `HTTP` | `HTTPS` | `TCP` | `UDP`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


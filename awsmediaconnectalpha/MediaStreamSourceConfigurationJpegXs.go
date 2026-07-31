package awsmediaconnectalpha


// Options for Media Stream Source Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encoding Encoding
//   var mediaStream MediaStream
//   var networkInterface NetworkInterface
//   var role Role
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   mediaStreamSourceConfigurationJpegXs := &MediaStreamSourceConfigurationJpegXs{
//   	Encoding: encoding,
//   	InputInterface: []VpcInterfaceConfig{
//   		&VpcInterfaceConfig{
//   			Name: jsii.String("name"),
//   			Role: role,
//   			SecurityGroups: []ISecurityGroup{
//   				securityGroup,
//   			},
//   			Subnet: subnet,
//
//   			// the properties below are optional
//   			NetworkInterfaceIds: []*string{
//   				jsii.String("networkInterfaceIds"),
//   			},
//   			NetworkInterfaceType: networkInterface,
//   		},
//   	},
//   	MediaStream: mediaStream,
//   	Port: jsii.Number(123),
//   }
//
// Experimental.
type MediaStreamSourceConfigurationJpegXs struct {
	// The format that was used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291.
	// If the encoding name is smpte291, set the color space to one of the following:
	// BT601, BT709, BT2020, or BT2100. For all other ancillary data streams, set the color space to SDR-NOCOLOR.
	// Experimental.
	Encoding Encoding `field:"required" json:"encoding" yaml:"encoding"`
	// The VPC interfaces where the media stream comes in from.
	// Experimental.
	InputInterface *[]*VpcInterfaceConfig `field:"required" json:"inputInterface" yaml:"inputInterface"`
	// The name of the media stream.
	// Experimental.
	MediaStream MediaStream `field:"required" json:"mediaStream" yaml:"mediaStream"`
	// The port that the flow listens on for this incoming media stream.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}


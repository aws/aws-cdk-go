package awsmedialive


// The settings for a PUSH type input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inputDestinationProperty := &InputDestinationProperty{
//   	Ip: jsii.String("ip"),
//   	Network: jsii.String("network"),
//   	NetworkRoutes: []interface{}{
//   		&InputDestinationRouteProperty{
//   			Cidr: jsii.String("cidr"),
//   			Gateway: jsii.String("gateway"),
//   		},
//   	},
//   	Port: jsii.String("port"),
//   	StaticIpAddress: jsii.String("staticIpAddress"),
//   	StreamName: jsii.String("streamName"),
//   	Url: jsii.String("url"),
//   	Vpc: &InputDestinationVpcProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html
//
type CfnInputPropsMixin_InputDestinationProperty struct {
	// The system-generated static IP address of endpoint.
	//
	// It remains fixed for the lifetime of the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-ip
	//
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The ID of the attached network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-network
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// If the push input has an input location of ON-PREM it's a requirement to specify what the route of the input is going to be on the customer local network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-networkroutes
	//
	NetworkRoutes interface{} `field:"optional" json:"networkRoutes" yaml:"networkRoutes"`
	// The port number for the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The static IP address for the input destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-staticipaddress
	//
	StaticIpAddress *string `field:"optional" json:"staticIpAddress" yaml:"staticIpAddress"`
	// The stream name for the input destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// This represents the endpoint that the customer stream will be pushed to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The properties for a VPC type input destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestination.html#cfn-medialive-input-inputdestination-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}


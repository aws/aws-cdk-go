package awsiotsecuretunneling


// The destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   destinationConfigProperty := &DestinationConfigProperty{
//   	Services: []*string{
//   		jsii.String("services"),
//   	},
//   	ThingName: jsii.String("thingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsecuretunneling-tunnel-destinationconfig.html
//
type CfnTunnelPropsMixin_DestinationConfigProperty struct {
	// A list of service names that identify the target application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsecuretunneling-tunnel-destinationconfig.html#cfn-iotsecuretunneling-tunnel-destinationconfig-services
	//
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// The name of the IoT thing to which you want to connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsecuretunneling-tunnel-destinationconfig.html#cfn-iotsecuretunneling-tunnel-destinationconfig-thingname
	//
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}


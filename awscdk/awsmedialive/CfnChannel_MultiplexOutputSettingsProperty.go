package awsmedialive


// Configuration of a Multiplex output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexOutputSettingsProperty := &MultiplexOutputSettingsProperty{
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexoutputsettings.html
//
type CfnChannel_MultiplexOutputSettingsProperty struct {
	// Destination is a Multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexoutputsettings.html#cfn-medialive-channel-multiplexoutputsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}


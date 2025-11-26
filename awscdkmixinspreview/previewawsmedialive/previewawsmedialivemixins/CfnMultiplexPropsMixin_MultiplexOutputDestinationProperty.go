package previewawsmedialivemixins


// Multiplex output destination settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexOutputDestinationProperty := &MultiplexOutputDestinationProperty{
//   	MultiplexMediaConnectOutputDestinationSettings: &MultiplexMediaConnectOutputDestinationSettingsProperty{
//   		EntitlementArn: jsii.String("entitlementArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexoutputdestination.html
//
type CfnMultiplexPropsMixin_MultiplexOutputDestinationProperty struct {
	// Multiplex MediaConnect output destination settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexoutputdestination.html#cfn-medialive-multiplex-multiplexoutputdestination-multiplexmediaconnectoutputdestinationsettings
	//
	MultiplexMediaConnectOutputDestinationSettings interface{} `field:"optional" json:"multiplexMediaConnectOutputDestinationSettings" yaml:"multiplexMediaConnectOutputDestinationSettings"`
}


package mixinsawsmedialive


// Multiplex MediaConnect output destination settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexMediaConnectOutputDestinationSettingsProperty := &MultiplexMediaConnectOutputDestinationSettingsProperty{
//   	EntitlementArn: jsii.String("entitlementArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexmediaconnectoutputdestinationsettings.html
//
type CfnMultiplexPropsMixin_MultiplexMediaConnectOutputDestinationSettingsProperty struct {
	// The MediaConnect entitlement ARN available as a Flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexmediaconnectoutputdestinationsettings.html#cfn-medialive-multiplex-multiplexmediaconnectoutputdestinationsettings-entitlementarn
	//
	EntitlementArn *string `field:"optional" json:"entitlementArn" yaml:"entitlementArn"`
}


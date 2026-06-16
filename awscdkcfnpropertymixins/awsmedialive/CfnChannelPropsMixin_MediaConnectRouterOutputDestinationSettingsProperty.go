package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mediaConnectRouterOutputDestinationSettingsProperty := &MediaConnectRouterOutputDestinationSettingsProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectrouteroutputdestinationsettings.html
//
type CfnChannelPropsMixin_MediaConnectRouterOutputDestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectrouteroutputdestinationsettings.html#cfn-medialive-channel-mediaconnectrouteroutputdestinationsettings-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectrouteroutputdestinationsettings.html#cfn-medialive-channel-mediaconnectrouteroutputdestinationsettings-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


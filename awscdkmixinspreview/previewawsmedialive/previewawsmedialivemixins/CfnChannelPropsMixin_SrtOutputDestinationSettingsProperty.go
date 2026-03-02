package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtOutputDestinationSettingsProperty := &SrtOutputDestinationSettingsProperty{
//   	ConnectionMode: jsii.String("connectionMode"),
//   	EncryptionPassphraseSecretArn: jsii.String("encryptionPassphraseSecretArn"),
//   	ListenerPort: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html
//
type CfnChannelPropsMixin_SrtOutputDestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-connectionmode
	//
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-encryptionpassphrasesecretarn
	//
	EncryptionPassphraseSecretArn *string `field:"optional" json:"encryptionPassphraseSecretArn" yaml:"encryptionPassphraseSecretArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-listenerport
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}


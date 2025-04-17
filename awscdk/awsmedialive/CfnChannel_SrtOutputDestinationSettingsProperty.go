package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtOutputDestinationSettingsProperty := &SrtOutputDestinationSettingsProperty{
//   	EncryptionPassphraseSecretArn: jsii.String("encryptionPassphraseSecretArn"),
//   	StreamId: jsii.String("streamId"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html
//
type CfnChannel_SrtOutputDestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-encryptionpassphrasesecretarn
	//
	EncryptionPassphraseSecretArn *string `field:"optional" json:"encryptionPassphraseSecretArn" yaml:"encryptionPassphraseSecretArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputdestinationsettings.html#cfn-medialive-channel-srtoutputdestinationsettings-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}


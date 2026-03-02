package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtListenerSettingsRequestProperty := &SrtListenerSettingsRequestProperty{
//   	Decryption: &SrtListenerDecryptionRequestProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   	},
//   	MinimumLatency: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettingsrequest.html
//
type CfnInputPropsMixin_SrtListenerSettingsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettingsrequest.html#cfn-medialive-input-srtlistenersettingsrequest-decryption
	//
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettingsrequest.html#cfn-medialive-input-srtlistenersettingsrequest-minimumlatency
	//
	MinimumLatency *float64 `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenersettingsrequest.html#cfn-medialive-input-srtlistenersettingsrequest-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}


package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtCallerSourceRequestProperty := &SrtCallerSourceRequestProperty{
//   	Decryption: &SrtCallerDecryptionRequestProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   	},
//   	MinimumLatency: jsii.Number(123),
//   	SrtListenerAddress: jsii.String("srtListenerAddress"),
//   	SrtListenerPort: jsii.String("srtListenerPort"),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersourcerequest.html
//
type CfnInputPropsMixin_SrtCallerSourceRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersourcerequest.html#cfn-medialive-input-srtcallersourcerequest-decryption
	//
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersourcerequest.html#cfn-medialive-input-srtcallersourcerequest-minimumlatency
	//
	MinimumLatency *float64 `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersourcerequest.html#cfn-medialive-input-srtcallersourcerequest-srtlisteneraddress
	//
	SrtListenerAddress *string `field:"optional" json:"srtListenerAddress" yaml:"srtListenerAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersourcerequest.html#cfn-medialive-input-srtcallersourcerequest-srtlistenerport
	//
	SrtListenerPort *string `field:"optional" json:"srtListenerPort" yaml:"srtListenerPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersourcerequest.html#cfn-medialive-input-srtcallersourcerequest-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}


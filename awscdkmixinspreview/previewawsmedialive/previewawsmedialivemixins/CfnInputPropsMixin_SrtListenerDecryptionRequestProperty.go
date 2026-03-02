package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   srtListenerDecryptionRequestProperty := &SrtListenerDecryptionRequestProperty{
//   	Algorithm: jsii.String("algorithm"),
//   	PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenerdecryptionrequest.html
//
type CfnInputPropsMixin_SrtListenerDecryptionRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenerdecryptionrequest.html#cfn-medialive-input-srtlistenerdecryptionrequest-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenerdecryptionrequest.html#cfn-medialive-input-srtlistenerdecryptionrequest-passphrasesecretarn
	//
	PassphraseSecretArn *string `field:"optional" json:"passphraseSecretArn" yaml:"passphraseSecretArn"`
}


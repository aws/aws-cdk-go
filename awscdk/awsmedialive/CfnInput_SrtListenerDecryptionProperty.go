package awsmedialive


// The decryption settings for the SRT listener source.
//
// Present only if the source has decryption enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtListenerDecryptionProperty := &SrtListenerDecryptionProperty{
//   	Algorithm: jsii.String("algorithm"),
//   	PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenerdecryption.html
//
type CfnInput_SrtListenerDecryptionProperty struct {
	// Placeholder documentation for Algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenerdecryption.html#cfn-medialive-input-srtlistenerdecryption-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The ARN for the secret in Secrets Manager.
	//
	// Someone in your organization must create a secret and provide you with its ARN. The secret holds the passphrase that MediaLive uses to decrypt the source content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtlistenerdecryption.html#cfn-medialive-input-srtlistenerdecryption-passphrasesecretarn
	//
	PassphraseSecretArn *string `field:"optional" json:"passphraseSecretArn" yaml:"passphraseSecretArn"`
}


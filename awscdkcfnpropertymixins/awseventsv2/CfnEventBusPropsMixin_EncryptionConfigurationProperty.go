package awseventsv2


// Encryption configuration for an event bus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventbus-encryptionconfiguration.html
//
type CfnEventBusPropsMixin_EncryptionConfigurationProperty struct {
	// The identifier of the AWS KMS customer managed key that the event bus uses to encrypt events.
	//
	// You can specify the key ARN, key ID, alias name, or alias ARN. If you do not specify a key, EventBridge uses an AWS owned key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventbus-encryptionconfiguration.html#cfn-eventsv2-eventbus-encryptionconfiguration-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
}


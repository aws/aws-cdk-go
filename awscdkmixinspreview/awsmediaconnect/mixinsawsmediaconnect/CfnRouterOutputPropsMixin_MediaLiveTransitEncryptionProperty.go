package mixinsawsmediaconnect


// The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive.
//
// This configuration determines whether encryption keys are automatically managed by the service or manually managed through AWS Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//
//   mediaLiveTransitEncryptionProperty := &MediaLiveTransitEncryptionProperty{
//   	EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
//   		Automatic: automatic,
//   		SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	EncryptionKeyType: jsii.String("encryptionKeyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialivetransitencryption.html
//
type CfnRouterOutputPropsMixin_MediaLiveTransitEncryptionProperty struct {
	// Configuration settings for the MediaLive transit encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialivetransitencryption.html#cfn-mediaconnect-routeroutput-medialivetransitencryption-encryptionkeyconfiguration
	//
	EncryptionKeyConfiguration interface{} `field:"optional" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialivetransitencryption.html#cfn-mediaconnect-routeroutput-medialivetransitencryption-encryptionkeytype
	//
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}


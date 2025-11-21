package awsmediaconnect


// The transit encryption settings for a router input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   routerInputTransitEncryptionProperty := &RouterInputTransitEncryptionProperty{
//   	EncryptionKeyConfiguration: &RouterInputTransitEncryptionKeyConfigurationProperty{
//   		Automatic: automatic,
//   		SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	EncryptionKeyType: jsii.String("encryptionKeyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputtransitencryption.html
//
type CfnRouterInput_RouterInputTransitEncryptionProperty struct {
	// Defines the configuration settings for transit encryption keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputtransitencryption.html#cfn-mediaconnect-routerinput-routerinputtransitencryption-encryptionkeyconfiguration
	//
	EncryptionKeyConfiguration interface{} `field:"required" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routerinputtransitencryption.html#cfn-mediaconnect-routerinput-routerinputtransitencryption-encryptionkeytype
	//
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}


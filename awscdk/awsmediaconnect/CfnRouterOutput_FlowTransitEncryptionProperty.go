package awsmediaconnect


// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   flowTransitEncryptionProperty := &FlowTransitEncryptionProperty{
//   	EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-flowtransitencryption.html
//
type CfnRouterOutput_FlowTransitEncryptionProperty struct {
	// Configuration settings for flow transit encryption keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-flowtransitencryption.html#cfn-mediaconnect-routeroutput-flowtransitencryption-encryptionkeyconfiguration
	//
	EncryptionKeyConfiguration interface{} `field:"required" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-flowtransitencryption.html#cfn-mediaconnect-routeroutput-flowtransitencryption-encryptionkeytype
	//
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}


package mixinsawsmediaconnect


// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   	EncryptionKeyType: jsii.String("encryptionKeyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-flowtransitencryption.html
//
type CfnRouterOutputPropsMixin_FlowTransitEncryptionProperty struct {
	// Configuration settings for flow transit encryption keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-flowtransitencryption.html#cfn-mediaconnect-routeroutput-flowtransitencryption-encryptionkeyconfiguration
	//
	EncryptionKeyConfiguration interface{} `field:"optional" json:"encryptionKeyConfiguration" yaml:"encryptionKeyConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-flowtransitencryption.html#cfn-mediaconnect-routeroutput-flowtransitencryption-encryptionkeytype
	//
	EncryptionKeyType *string `field:"optional" json:"encryptionKeyType" yaml:"encryptionKeyType"`
}


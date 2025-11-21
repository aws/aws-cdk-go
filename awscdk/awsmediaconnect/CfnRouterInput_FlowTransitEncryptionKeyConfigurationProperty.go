package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   flowTransitEncryptionKeyConfigurationProperty := &FlowTransitEncryptionKeyConfigurationProperty{
//   	Automatic: automatic,
//   	SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-flowtransitencryptionkeyconfiguration.html
//
type CfnRouterInput_FlowTransitEncryptionKeyConfigurationProperty struct {
	// Configuration settings for automatic encryption key management, where MediaConnect handles key creation and rotation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-flowtransitencryptionkeyconfiguration.html#cfn-mediaconnect-routerinput-flowtransitencryptionkeyconfiguration-automatic
	//
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-flowtransitencryptionkeyconfiguration.html#cfn-mediaconnect-routerinput-flowtransitencryptionkeyconfiguration-secretsmanager
	//
	SecretsManager interface{} `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}


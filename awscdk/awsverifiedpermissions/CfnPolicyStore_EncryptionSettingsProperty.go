package awsverifiedpermissions


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var default_ interface{}
//
//   encryptionSettingsProperty := &EncryptionSettingsProperty{
//   	Default: default_,
//   	KmsEncryptionSettings: &KmsEncryptionSettingsProperty{
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		EncryptionContext: map[string]*string{
//   			"encryptionContextKey": jsii.String("encryptionContext"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-encryptionsettings.html
//
type CfnPolicyStore_EncryptionSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-encryptionsettings.html#cfn-verifiedpermissions-policystore-encryptionsettings-default
	//
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-encryptionsettings.html#cfn-verifiedpermissions-policystore-encryptionsettings-kmsencryptionsettings
	//
	KmsEncryptionSettings interface{} `field:"optional" json:"kmsEncryptionSettings" yaml:"kmsEncryptionSettings"`
}


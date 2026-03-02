package awsverifiedpermissions


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kmsEncryptionSettingsProperty := &KmsEncryptionSettingsProperty{
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	EncryptionContext: map[string]*string{
//   		"encryptionContextKey": jsii.String("encryptionContext"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-kmsencryptionsettings.html
//
type CfnPolicyStore_KmsEncryptionSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-kmsencryptionsettings.html#cfn-verifiedpermissions-policystore-kmsencryptionsettings-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-kmsencryptionsettings.html#cfn-verifiedpermissions-policystore-kmsencryptionsettings-encryptioncontext
	//
	EncryptionContext interface{} `field:"optional" json:"encryptionContext" yaml:"encryptionContext"`
}


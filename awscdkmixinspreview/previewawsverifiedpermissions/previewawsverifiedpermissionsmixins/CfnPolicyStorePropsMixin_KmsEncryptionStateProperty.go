package previewawsverifiedpermissionsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kmsEncryptionStateProperty := &KmsEncryptionStateProperty{
//   	EncryptionContext: map[string]*string{
//   		"encryptionContextKey": jsii.String("encryptionContext"),
//   	},
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-kmsencryptionstate.html
//
type CfnPolicyStorePropsMixin_KmsEncryptionStateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-kmsencryptionstate.html#cfn-verifiedpermissions-policystore-kmsencryptionstate-encryptioncontext
	//
	EncryptionContext interface{} `field:"optional" json:"encryptionContext" yaml:"encryptionContext"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-kmsencryptionstate.html#cfn-verifiedpermissions-policystore-kmsencryptionstate-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}


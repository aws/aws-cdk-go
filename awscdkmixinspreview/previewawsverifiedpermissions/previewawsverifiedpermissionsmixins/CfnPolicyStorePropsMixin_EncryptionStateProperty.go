package previewawsverifiedpermissionsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var default_ interface{}
//
//   encryptionStateProperty := &EncryptionStateProperty{
//   	Default: default_,
//   	KmsEncryptionState: &KmsEncryptionStateProperty{
//   		EncryptionContext: map[string]*string{
//   			"encryptionContextKey": jsii.String("encryptionContext"),
//   		},
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-encryptionstate.html
//
type CfnPolicyStorePropsMixin_EncryptionStateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-encryptionstate.html#cfn-verifiedpermissions-policystore-encryptionstate-default
	//
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-encryptionstate.html#cfn-verifiedpermissions-policystore-encryptionstate-kmsencryptionstate
	//
	KmsEncryptionState interface{} `field:"optional" json:"kmsEncryptionState" yaml:"kmsEncryptionState"`
}


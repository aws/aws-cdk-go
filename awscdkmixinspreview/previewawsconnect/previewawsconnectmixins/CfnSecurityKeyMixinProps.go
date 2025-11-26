package previewawsconnectmixins


// Properties for CfnSecurityKeyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityKeyMixinProps := &CfnSecurityKeyMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securitykey.html
//
type CfnSecurityKeyMixinProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securitykey.html#cfn-connect-securitykey-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// A valid security key in PEM format. For example:.
	//
	// `"-----BEGIN PUBLIC KEY-----\ [a lot of characters] ----END PUBLIC KEY-----"`
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securitykey.html#cfn-connect-securitykey-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}


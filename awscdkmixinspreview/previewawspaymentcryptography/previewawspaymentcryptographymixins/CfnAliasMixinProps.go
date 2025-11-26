package previewawspaymentcryptographymixins


// Properties for CfnAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAliasMixinProps := &CfnAliasMixinProps{
//   	AliasName: jsii.String("aliasName"),
//   	KeyArn: jsii.String("keyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-alias.html
//
type CfnAliasMixinProps struct {
	// A friendly name that you can use to refer to a key. The value must begin with `alias/` .
	//
	// > Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-alias.html#cfn-paymentcryptography-alias-aliasname
	//
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// The `KeyARN` of the key associated with the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-paymentcryptography-alias.html#cfn-paymentcryptography-alias-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}


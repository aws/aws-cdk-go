package awscloudfront


// Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicKeyConfigProperty := &PublicKeyConfigProperty{
//   	CallerReference: jsii.String("callerReference"),
//   	EncodedKey: jsii.String("encodedKey"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-publickey-publickeyconfig.html
//
type CfnPublicKey_PublicKeyConfigProperty struct {
	// A string included in the request to help make sure that the request can't be replayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-publickey-publickeyconfig.html#cfn-cloudfront-publickey-publickeyconfig-callerreference
	//
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// The public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-publickey-publickeyconfig.html#cfn-cloudfront-publickey-publickeyconfig-encodedkey
	//
	EncodedKey *string `field:"required" json:"encodedKey" yaml:"encodedKey"`
	// A name to help identify the public key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-publickey-publickeyconfig.html#cfn-cloudfront-publickey-publickeyconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the public key.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-publickey-publickeyconfig.html#cfn-cloudfront-publickey-publickeyconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}


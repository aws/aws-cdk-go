package awscloudfront


// Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicKeyConfigProperty := &publicKeyConfigProperty{
//   	callerReference: jsii.String("callerReference"),
//   	encodedKey: jsii.String("encodedKey"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   }
//
type CfnPublicKey_PublicKeyConfigProperty struct {
	// A string included in the request to help make sure that the request can’t be replayed.
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// The public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
	EncodedKey *string `field:"required" json:"encodedKey" yaml:"encodedKey"`
	// A name to help identify the public key.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the public key.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}


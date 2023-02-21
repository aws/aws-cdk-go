package awscloudfront


// Properties for defining a `CfnPublicKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPublicKeyProps := &CfnPublicKeyProps{
//   	PublicKeyConfig: &PublicKeyConfigProperty{
//   		CallerReference: jsii.String("callerReference"),
//   		EncodedKey: jsii.String("encodedKey"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Comment: jsii.String("comment"),
//   	},
//   }
//
type CfnPublicKeyProps struct {
	// Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
	PublicKeyConfig interface{} `field:"required" json:"publicKeyConfig" yaml:"publicKeyConfig"`
}


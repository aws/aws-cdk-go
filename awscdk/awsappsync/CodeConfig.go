package awsappsync


// Result of binding `Code` into a `Function`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfig := &CodeConfig{
//   	InlineCode: jsii.String("inlineCode"),
//   	S3Location: jsii.String("s3Location"),
//   }
//
type CodeConfig struct {
	// Inline code (mutually exclusive with `s3Location`).
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode`.
	S3Location *string `field:"optional" json:"s3Location" yaml:"s3Location"`
}


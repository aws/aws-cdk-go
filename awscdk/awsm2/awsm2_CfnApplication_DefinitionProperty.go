package awsm2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionProperty := &definitionProperty{
//   	content: jsii.String("content"),
//   	s3Location: jsii.String("s3Location"),
//   }
//
type CfnApplication_DefinitionProperty struct {
	// `CfnApplication.DefinitionProperty.Content`.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// `CfnApplication.DefinitionProperty.S3Location`.
	S3Location *string `field:"optional" json:"s3Location" yaml:"s3Location"`
}


package awsm2


// The application definition for a particular application.
//
// You can specify either inline JSON or an Amazon S3 bucket location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionProperty := &DefinitionProperty{
//   	Content: jsii.String("content"),
//   	S3Location: jsii.String("s3Location"),
//   }
//
type CfnApplication_DefinitionProperty struct {
	// The content of the application definition.
	//
	// This is a JSON object that contains the resource configuration/definitions that identify an application.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The S3 bucket that contains the application definition.
	S3Location *string `field:"optional" json:"s3Location" yaml:"s3Location"`
}


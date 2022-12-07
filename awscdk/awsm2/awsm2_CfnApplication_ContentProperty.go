package awsm2


// The content of the application definition.
//
// This is a JSON object that contains the resource configuration/definitions that identify an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentProperty := &contentProperty{
//   	s3Location: jsii.String("s3Location"),
//   }
//
type CfnApplication_ContentProperty struct {
	// The content of the application definition.
	//
	// This is a JSON object that contains the resource configuration/definitions that identify an application.
	S3Location *string `field:"required" json:"s3Location" yaml:"s3Location"`
}


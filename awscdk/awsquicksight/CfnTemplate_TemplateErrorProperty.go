package awsquicksight


// List of errors that occurred when the template version creation failed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateErrorProperty := &TemplateErrorProperty{
//   	Message: jsii.String("message"),
//   	Type: jsii.String("type"),
//   }
//
type CfnTemplate_TemplateErrorProperty struct {
	// Description of the error type.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Type of error.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


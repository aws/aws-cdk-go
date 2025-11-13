package interfacesawspinpoint


// A reference to a EmailTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailTemplateReference := &EmailTemplateReference{
//   	EmailTemplateArn: jsii.String("emailTemplateArn"),
//   	EmailTemplateId: jsii.String("emailTemplateId"),
//   }
//
type EmailTemplateReference struct {
	// The ARN of the EmailTemplate resource.
	EmailTemplateArn *string `field:"required" json:"emailTemplateArn" yaml:"emailTemplateArn"`
	// The Id of the EmailTemplate resource.
	EmailTemplateId *string `field:"required" json:"emailTemplateId" yaml:"emailTemplateId"`
}


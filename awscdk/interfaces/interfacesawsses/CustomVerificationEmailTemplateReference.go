package interfacesawsses


// A reference to a CustomVerificationEmailTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customVerificationEmailTemplateReference := &CustomVerificationEmailTemplateReference{
//   	TemplateName: jsii.String("templateName"),
//   }
//
type CustomVerificationEmailTemplateReference struct {
	// The TemplateName of the CustomVerificationEmailTemplate resource.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}


package awsiot


// A reference to a ProvisioningTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisioningTemplateReference := &ProvisioningTemplateReference{
//   	TemplateName: jsii.String("templateName"),
//   }
//
type ProvisioningTemplateReference struct {
	// The TemplateName of the ProvisioningTemplate resource.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}


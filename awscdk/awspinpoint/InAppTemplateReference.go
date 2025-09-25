package awspinpoint


// A reference to a InAppTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppTemplateReference := &InAppTemplateReference{
//   	InAppTemplateArn: jsii.String("inAppTemplateArn"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
type InAppTemplateReference struct {
	// The ARN of the InAppTemplate resource.
	InAppTemplateArn *string `field:"required" json:"inAppTemplateArn" yaml:"inAppTemplateArn"`
	// The TemplateName of the InAppTemplate resource.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}


package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateErrorProperty := &templateErrorProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnTemplate_TemplateErrorProperty struct {
	// `CfnTemplate.TemplateErrorProperty.Message`.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// `CfnTemplate.TemplateErrorProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


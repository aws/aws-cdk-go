package interfacesawsses


// A reference to a Template resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateReference := &TemplateReference{
//   	TemplateId: jsii.String("templateId"),
//   }
//
type TemplateReference struct {
	// The Id of the Template resource.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
}


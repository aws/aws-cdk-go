package interfacesawscloudformation


// A reference to a GeneratedTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generatedTemplateReference := &GeneratedTemplateReference{
//   	GeneratedTemplateId: jsii.String("generatedTemplateId"),
//   }
//
type GeneratedTemplateReference struct {
	// The GeneratedTemplateId of the GeneratedTemplate resource.
	GeneratedTemplateId *string `field:"required" json:"generatedTemplateId" yaml:"generatedTemplateId"`
}


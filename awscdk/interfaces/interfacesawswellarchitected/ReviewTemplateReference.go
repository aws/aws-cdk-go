package interfacesawswellarchitected


// A reference to a ReviewTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reviewTemplateReference := &ReviewTemplateReference{
//   	TemplateArn: jsii.String("templateArn"),
//   }
//
type ReviewTemplateReference struct {
	// The TemplateArn of the ReviewTemplate resource.
	TemplateArn *string `field:"required" json:"templateArn" yaml:"templateArn"`
}


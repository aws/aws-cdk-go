package interfacesawsiotfleetwise


// A reference to a StateTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateTemplateReference := &StateTemplateReference{
//   	StateTemplateArn: jsii.String("stateTemplateArn"),
//   	StateTemplateName: jsii.String("stateTemplateName"),
//   }
//
type StateTemplateReference struct {
	// The ARN of the StateTemplate resource.
	StateTemplateArn *string `field:"required" json:"stateTemplateArn" yaml:"stateTemplateArn"`
	// The Name of the StateTemplate resource.
	StateTemplateName *string `field:"required" json:"stateTemplateName" yaml:"stateTemplateName"`
}


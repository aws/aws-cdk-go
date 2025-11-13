package interfacesawspinpoint


// A reference to a PushTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pushTemplateReference := &PushTemplateReference{
//   	PushTemplateArn: jsii.String("pushTemplateArn"),
//   	PushTemplateId: jsii.String("pushTemplateId"),
//   }
//
type PushTemplateReference struct {
	// The ARN of the PushTemplate resource.
	PushTemplateArn *string `field:"required" json:"pushTemplateArn" yaml:"pushTemplateArn"`
	// The Id of the PushTemplate resource.
	PushTemplateId *string `field:"required" json:"pushTemplateId" yaml:"pushTemplateId"`
}


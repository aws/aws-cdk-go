package interfacesawswisdom


// A reference to a MessageTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageTemplateReference := &MessageTemplateReference{
//   	MessageTemplateArn: jsii.String("messageTemplateArn"),
//   }
//
type MessageTemplateReference struct {
	// The MessageTemplateArn of the MessageTemplate resource.
	MessageTemplateArn *string `field:"required" json:"messageTemplateArn" yaml:"messageTemplateArn"`
}


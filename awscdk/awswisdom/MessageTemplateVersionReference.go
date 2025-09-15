package awswisdom


// A reference to a MessageTemplateVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageTemplateVersionReference := &MessageTemplateVersionReference{
//   	MessageTemplateVersionArn: jsii.String("messageTemplateVersionArn"),
//   }
//
type MessageTemplateVersionReference struct {
	// The MessageTemplateVersionArn of the MessageTemplateVersion resource.
	MessageTemplateVersionArn *string `field:"required" json:"messageTemplateVersionArn" yaml:"messageTemplateVersionArn"`
}


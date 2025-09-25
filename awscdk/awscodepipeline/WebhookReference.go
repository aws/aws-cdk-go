package awscodepipeline


// A reference to a Webhook resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookReference := &WebhookReference{
//   	WebhookId: jsii.String("webhookId"),
//   }
//
type WebhookReference struct {
	// The Id of the Webhook resource.
	WebhookId *string `field:"required" json:"webhookId" yaml:"webhookId"`
}


package interfacesawscodepipeline


// A reference to a Webhook resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookReference := &WebhookReference{
//   	WebhookName: jsii.String("webhookName"),
//   }
//
type WebhookReference struct {
	// The Name of the Webhook resource.
	WebhookName *string `field:"required" json:"webhookName" yaml:"webhookName"`
}


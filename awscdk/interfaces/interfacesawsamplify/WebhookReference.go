package interfacesawsamplify


// A reference to a Webhook resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookReference := &WebhookReference{
//   	WebhookArn: jsii.String("webhookArn"),
//   }
//
type WebhookReference struct {
	// The Arn of the Webhook resource.
	WebhookArn *string `field:"required" json:"webhookArn" yaml:"webhookArn"`
}


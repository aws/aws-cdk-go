package awssecuritylake


// A reference to a Subscriber resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriberReference := &SubscriberReference{
//   	SubscriberArn: jsii.String("subscriberArn"),
//   }
//
type SubscriberReference struct {
	// The SubscriberArn of the Subscriber resource.
	SubscriberArn *string `field:"required" json:"subscriberArn" yaml:"subscriberArn"`
}


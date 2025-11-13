package interfacesawsvpclattice


// A reference to a AccessLogSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSubscriptionReference := &AccessLogSubscriptionReference{
//   	AccessLogSubscriptionArn: jsii.String("accessLogSubscriptionArn"),
//   }
//
type AccessLogSubscriptionReference struct {
	// The Arn of the AccessLogSubscription resource.
	AccessLogSubscriptionArn *string `field:"required" json:"accessLogSubscriptionArn" yaml:"accessLogSubscriptionArn"`
}


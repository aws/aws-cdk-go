package awssecurityhub


// A reference to a ProductSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   productSubscriptionReference := &ProductSubscriptionReference{
//   	ProductSubscriptionArn: jsii.String("productSubscriptionArn"),
//   }
//
type ProductSubscriptionReference struct {
	// The ProductSubscriptionArn of the ProductSubscription resource.
	ProductSubscriptionArn *string `field:"required" json:"productSubscriptionArn" yaml:"productSubscriptionArn"`
}


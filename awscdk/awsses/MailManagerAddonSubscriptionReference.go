package awsses


// A reference to a MailManagerAddonSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerAddonSubscriptionReference := &MailManagerAddonSubscriptionReference{
//   	AddonSubscriptionId: jsii.String("addonSubscriptionId"),
//   }
//
type MailManagerAddonSubscriptionReference struct {
	// The AddonSubscriptionId of the MailManagerAddonSubscription resource.
	AddonSubscriptionId *string `field:"required" json:"addonSubscriptionId" yaml:"addonSubscriptionId"`
}


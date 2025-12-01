package interfacesawsses


// A reference to a MailManagerAddonSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerAddonSubscriptionReference := &MailManagerAddonSubscriptionReference{
//   	AddonSubscriptionArn: jsii.String("addonSubscriptionArn"),
//   	AddonSubscriptionId: jsii.String("addonSubscriptionId"),
//   }
//
type MailManagerAddonSubscriptionReference struct {
	// The ARN of the MailManagerAddonSubscription resource.
	AddonSubscriptionArn *string `field:"required" json:"addonSubscriptionArn" yaml:"addonSubscriptionArn"`
	// The AddonSubscriptionId of the MailManagerAddonSubscription resource.
	AddonSubscriptionId *string `field:"required" json:"addonSubscriptionId" yaml:"addonSubscriptionId"`
}


package interfacesawsses


// A reference to a MailManagerAddonInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerAddonInstanceReference := &MailManagerAddonInstanceReference{
//   	AddonInstanceId: jsii.String("addonInstanceId"),
//   }
//
type MailManagerAddonInstanceReference struct {
	// The AddonInstanceId of the MailManagerAddonInstance resource.
	AddonInstanceId *string `field:"required" json:"addonInstanceId" yaml:"addonInstanceId"`
}


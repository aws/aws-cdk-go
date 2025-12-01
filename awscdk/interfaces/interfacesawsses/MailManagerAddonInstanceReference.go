package interfacesawsses


// A reference to a MailManagerAddonInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerAddonInstanceReference := &MailManagerAddonInstanceReference{
//   	AddonInstanceArn: jsii.String("addonInstanceArn"),
//   	AddonInstanceId: jsii.String("addonInstanceId"),
//   }
//
type MailManagerAddonInstanceReference struct {
	// The ARN of the MailManagerAddonInstance resource.
	AddonInstanceArn *string `field:"required" json:"addonInstanceArn" yaml:"addonInstanceArn"`
	// The AddonInstanceId of the MailManagerAddonInstance resource.
	AddonInstanceId *string `field:"required" json:"addonInstanceId" yaml:"addonInstanceId"`
}


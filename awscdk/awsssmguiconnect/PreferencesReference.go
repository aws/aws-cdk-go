package awsssmguiconnect


// A reference to a Preferences resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preferencesReference := &PreferencesReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type PreferencesReference struct {
	// The AccountId of the Preferences resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}


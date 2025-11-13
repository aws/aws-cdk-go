package interfacesawsworkspacesweb


// A reference to a UserSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userSettingsReference := &UserSettingsReference{
//   	UserSettingsArn: jsii.String("userSettingsArn"),
//   }
//
type UserSettingsReference struct {
	// The UserSettingsArn of the UserSettings resource.
	UserSettingsArn *string `field:"required" json:"userSettingsArn" yaml:"userSettingsArn"`
}


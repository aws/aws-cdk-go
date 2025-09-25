package awsworkspacesweb


// A reference to a UserAccessLoggingSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userAccessLoggingSettingsReference := &UserAccessLoggingSettingsReference{
//   	UserAccessLoggingSettingsArn: jsii.String("userAccessLoggingSettingsArn"),
//   }
//
type UserAccessLoggingSettingsReference struct {
	// The UserAccessLoggingSettingsArn of the UserAccessLoggingSettings resource.
	UserAccessLoggingSettingsArn *string `field:"required" json:"userAccessLoggingSettingsArn" yaml:"userAccessLoggingSettingsArn"`
}


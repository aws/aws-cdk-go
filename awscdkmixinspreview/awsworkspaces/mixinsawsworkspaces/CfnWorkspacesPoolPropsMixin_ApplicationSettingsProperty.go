package mixinsawsworkspaces


// The persistent application settings for users in the pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationSettingsProperty := &ApplicationSettingsProperty{
//   	SettingsGroup: jsii.String("settingsGroup"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html
//
type CfnWorkspacesPoolPropsMixin_ApplicationSettingsProperty struct {
	// The path prefix for the S3 bucket where users’ persistent application settings are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html#cfn-workspaces-workspacespool-applicationsettings-settingsgroup
	//
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
	// Enables or disables persistent application settings for users during their pool sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html#cfn-workspaces-workspacespool-applicationsettings-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}


package awsworkspaces


// The persistent application settings for users in the pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSettingsProperty := &ApplicationSettingsProperty{
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	SettingsGroup: jsii.String("settingsGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html
//
type CfnWorkspacesPool_ApplicationSettingsProperty struct {
	// Enables or disables persistent application settings for users during their pool sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html#cfn-workspaces-workspacespool-applicationsettings-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// The path prefix for the S3 bucket where users’ persistent application settings are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-applicationsettings.html#cfn-workspaces-workspacespool-applicationsettings-settingsgroup
	//
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
}


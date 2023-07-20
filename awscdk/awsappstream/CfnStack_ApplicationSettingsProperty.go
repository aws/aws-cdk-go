package awsappstream


// The persistent application settings for users of a stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSettingsProperty := &ApplicationSettingsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	SettingsGroup: jsii.String("settingsGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html
//
type CfnStack_ApplicationSettingsProperty struct {
	// Enables or disables persistent application settings for users during their streaming sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html#cfn-appstream-stack-applicationsettings-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The path prefix for the S3 bucket where users’ persistent application settings are stored.
	//
	// You can allow the same persistent application settings to be used across multiple stacks by specifying the same settings group for each stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html#cfn-appstream-stack-applicationsettings-settingsgroup
	//
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
}


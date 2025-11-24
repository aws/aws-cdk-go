package mixinsawsworkspacesweb


// The configuration of the toolbar.
//
// This allows administrators to select the toolbar type and visual mode, set maximum display resolution for sessions, and choose which items are visible to end users during their sessions. If administrators do not modify these settings, end users retain control over their toolbar preferences.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   toolbarConfigurationProperty := &ToolbarConfigurationProperty{
//   	HiddenToolbarItems: []*string{
//   		jsii.String("hiddenToolbarItems"),
//   	},
//   	MaxDisplayResolution: jsii.String("maxDisplayResolution"),
//   	ToolbarType: jsii.String("toolbarType"),
//   	VisualMode: jsii.String("visualMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-toolbarconfiguration.html
//
type CfnUserSettingsPropsMixin_ToolbarConfigurationProperty struct {
	// The list of toolbar items to be hidden.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-toolbarconfiguration.html#cfn-workspacesweb-usersettings-toolbarconfiguration-hiddentoolbaritems
	//
	HiddenToolbarItems *[]*string `field:"optional" json:"hiddenToolbarItems" yaml:"hiddenToolbarItems"`
	// The maximum display resolution that is allowed for the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-toolbarconfiguration.html#cfn-workspacesweb-usersettings-toolbarconfiguration-maxdisplayresolution
	//
	MaxDisplayResolution *string `field:"optional" json:"maxDisplayResolution" yaml:"maxDisplayResolution"`
	// The type of toolbar displayed during the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-toolbarconfiguration.html#cfn-workspacesweb-usersettings-toolbarconfiguration-toolbartype
	//
	ToolbarType *string `field:"optional" json:"toolbarType" yaml:"toolbarType"`
	// The visual mode of the toolbar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-toolbarconfiguration.html#cfn-workspacesweb-usersettings-toolbarconfiguration-visualmode
	//
	VisualMode *string `field:"optional" json:"visualMode" yaml:"visualMode"`
}


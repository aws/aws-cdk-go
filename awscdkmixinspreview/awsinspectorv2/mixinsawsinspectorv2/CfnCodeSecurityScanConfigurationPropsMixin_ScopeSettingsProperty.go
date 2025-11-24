package mixinsawsinspectorv2


// The scope settings that define which repositories will be scanned.
//
// If the `ScopeSetting` parameter is `ALL` the scan configuration applies to all existing and future projects imported into Amazon Inspector .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scopeSettingsProperty := &ScopeSettingsProperty{
//   	ProjectSelectionScope: jsii.String("projectSelectionScope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-scopesettings.html
//
type CfnCodeSecurityScanConfigurationPropsMixin_ScopeSettingsProperty struct {
	// The scope of projects to be selected for scanning within the integrated repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-scopesettings.html#cfn-inspectorv2-codesecurityscanconfiguration-scopesettings-projectselectionscope
	//
	ProjectSelectionScope *string `field:"optional" json:"projectSelectionScope" yaml:"projectSelectionScope"`
}


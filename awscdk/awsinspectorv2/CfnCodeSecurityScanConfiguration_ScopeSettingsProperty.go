package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeSettingsProperty := &ScopeSettingsProperty{
//   	ProjectSelectionScope: jsii.String("projectSelectionScope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-scopesettings.html
//
type CfnCodeSecurityScanConfiguration_ScopeSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-scopesettings.html#cfn-inspectorv2-codesecurityscanconfiguration-scopesettings-projectselectionscope
	//
	ProjectSelectionScope *string `field:"optional" json:"projectSelectionScope" yaml:"projectSelectionScope"`
}


package previewawsconnectmixins


// Contains primary color configuration for a workspace theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   palettePrimaryProperty := &PalettePrimaryProperty{
//   	Active: jsii.String("active"),
//   	ContrastText: jsii.String("contrastText"),
//   	Default: jsii.String("default"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteprimary.html
//
type CfnWorkspacePropsMixin_PalettePrimaryProperty struct {
	// The primary color used for active states.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteprimary.html#cfn-connect-workspace-paletteprimary-active
	//
	Active *string `field:"optional" json:"active" yaml:"active"`
	// The text color that contrasts with the primary color for readability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteprimary.html#cfn-connect-workspace-paletteprimary-contrasttext
	//
	ContrastText *string `field:"optional" json:"contrastText" yaml:"contrastText"`
	// The default primary color used throughout the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteprimary.html#cfn-connect-workspace-paletteprimary-default
	//
	Default *string `field:"optional" json:"default" yaml:"default"`
}


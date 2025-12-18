package previewawsconnectmixins


// Contains typography configuration for a workspace theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workspaceThemeTypographyProperty := &WorkspaceThemeTypographyProperty{
//   	FontFamily: &FontFamilyProperty{
//   		Default: jsii.String("default"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemetypography.html
//
type CfnWorkspacePropsMixin_WorkspaceThemeTypographyProperty struct {
	// The font family configuration for text in the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacethemetypography.html#cfn-connect-workspace-workspacethemetypography-fontfamily
	//
	FontFamily interface{} `field:"optional" json:"fontFamily" yaml:"fontFamily"`
}


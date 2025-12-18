package awsconnect


// Contains font family configuration for workspace themes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontFamilyProperty := &FontFamilyProperty{
//   	Default: jsii.String("default"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-fontfamily.html
//
type CfnWorkspace_FontFamilyProperty struct {
	// The default font family to use in the workspace theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-fontfamily.html#cfn-connect-workspace-fontfamily-default
	//
	Default *string `field:"optional" json:"default" yaml:"default"`
}


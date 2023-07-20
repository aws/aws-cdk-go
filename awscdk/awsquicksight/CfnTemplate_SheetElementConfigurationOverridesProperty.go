package awsquicksight


// The override configuration of the rendering rules of a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetElementConfigurationOverridesProperty := &SheetElementConfigurationOverridesProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetelementconfigurationoverrides.html
//
type CfnTemplate_SheetElementConfigurationOverridesProperty struct {
	// Determines whether or not the overrides are visible. Choose one of the following options:.
	//
	// - `VISIBLE`
	// - `HIDDEN`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetelementconfigurationoverrides.html#cfn-quicksight-template-sheetelementconfigurationoverrides-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}


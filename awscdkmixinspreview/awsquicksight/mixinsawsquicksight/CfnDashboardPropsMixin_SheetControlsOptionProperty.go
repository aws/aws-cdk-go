package mixinsawsquicksight


// Sheet controls option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetControlsOptionProperty := &SheetControlsOptionProperty{
//   	VisibilityState: jsii.String("visibilityState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolsoption.html
//
type CfnDashboardPropsMixin_SheetControlsOptionProperty struct {
	// Visibility state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolsoption.html#cfn-quicksight-dashboard-sheetcontrolsoption-visibilitystate
	//
	VisibilityState *string `field:"optional" json:"visibilityState" yaml:"visibilityState"`
}


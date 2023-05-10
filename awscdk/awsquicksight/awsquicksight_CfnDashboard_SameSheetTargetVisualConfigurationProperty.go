package awsquicksight


// The configuration of the same-sheet target visuals that you want to be filtered.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sameSheetTargetVisualConfigurationProperty := &SameSheetTargetVisualConfigurationProperty{
//   	TargetVisualOptions: jsii.String("targetVisualOptions"),
//   	TargetVisuals: []*string{
//   		jsii.String("targetVisuals"),
//   	},
//   }
//
type CfnDashboard_SameSheetTargetVisualConfigurationProperty struct {
	// The options that choose the target visual in the same sheet.
	//
	// Valid values are defined as follows:
	//
	// - `ALL_VISUALS` : Applies the filter operation to all visuals in the same sheet.
	TargetVisualOptions *string `field:"optional" json:"targetVisualOptions" yaml:"targetVisualOptions"`
	// A list of the target visual IDs that are located in the same sheet of the analysis.
	TargetVisuals *[]*string `field:"optional" json:"targetVisuals" yaml:"targetVisuals"`
}


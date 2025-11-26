package previewawsquicksightmixins


// The configuration of target visuals that you want to be filtered.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterOperationTargetVisualsConfigurationProperty := &FilterOperationTargetVisualsConfigurationProperty{
//   	SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   		TargetVisualOptions: jsii.String("targetVisualOptions"),
//   		TargetVisuals: []*string{
//   			jsii.String("targetVisuals"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filteroperationtargetvisualsconfiguration.html
//
type CfnTemplatePropsMixin_FilterOperationTargetVisualsConfigurationProperty struct {
	// The configuration of the same-sheet target visuals that you want to be filtered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filteroperationtargetvisualsconfiguration.html#cfn-quicksight-template-filteroperationtargetvisualsconfiguration-samesheettargetvisualconfiguration
	//
	SameSheetTargetVisualConfiguration interface{} `field:"optional" json:"sameSheetTargetVisualConfiguration" yaml:"sameSheetTargetVisualConfiguration"`
}


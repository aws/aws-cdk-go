package awsquicksight


// Determines the border options for a table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalTableBorderOptionsProperty := &GlobalTableBorderOptionsProperty{
//   	SideSpecificBorder: &TableSideBorderOptionsProperty{
//   		Bottom: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		InnerHorizontal: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		InnerVertical: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		Left: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		Right: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   		Top: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   	},
//   	UniformBorder: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-globaltableborderoptions.html
//
type CfnAnalysis_GlobalTableBorderOptionsProperty struct {
	// Determines the options for side specific border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-globaltableborderoptions.html#cfn-quicksight-analysis-globaltableborderoptions-sidespecificborder
	//
	SideSpecificBorder interface{} `field:"optional" json:"sideSpecificBorder" yaml:"sideSpecificBorder"`
	// Determines the options for uniform border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-globaltableborderoptions.html#cfn-quicksight-analysis-globaltableborderoptions-uniformborder
	//
	UniformBorder interface{} `field:"optional" json:"uniformBorder" yaml:"uniformBorder"`
}


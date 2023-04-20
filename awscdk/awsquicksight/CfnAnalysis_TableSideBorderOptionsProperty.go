package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableSideBorderOptionsProperty := &TableSideBorderOptionsProperty{
//   	Bottom: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	InnerHorizontal: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	InnerVertical: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	Left: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	Right: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	Top: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_TableSideBorderOptionsProperty struct {
	// `CfnAnalysis.TableSideBorderOptionsProperty.Bottom`.
	Bottom interface{} `field:"optional" json:"bottom" yaml:"bottom"`
	// `CfnAnalysis.TableSideBorderOptionsProperty.InnerHorizontal`.
	InnerHorizontal interface{} `field:"optional" json:"innerHorizontal" yaml:"innerHorizontal"`
	// `CfnAnalysis.TableSideBorderOptionsProperty.InnerVertical`.
	InnerVertical interface{} `field:"optional" json:"innerVertical" yaml:"innerVertical"`
	// `CfnAnalysis.TableSideBorderOptionsProperty.Left`.
	Left interface{} `field:"optional" json:"left" yaml:"left"`
	// `CfnAnalysis.TableSideBorderOptionsProperty.Right`.
	Right interface{} `field:"optional" json:"right" yaml:"right"`
	// `CfnAnalysis.TableSideBorderOptionsProperty.Top`.
	Top interface{} `field:"optional" json:"top" yaml:"top"`
}


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
type CfnTemplate_TableSideBorderOptionsProperty struct {
	// `CfnTemplate.TableSideBorderOptionsProperty.Bottom`.
	Bottom interface{} `field:"optional" json:"bottom" yaml:"bottom"`
	// `CfnTemplate.TableSideBorderOptionsProperty.InnerHorizontal`.
	InnerHorizontal interface{} `field:"optional" json:"innerHorizontal" yaml:"innerHorizontal"`
	// `CfnTemplate.TableSideBorderOptionsProperty.InnerVertical`.
	InnerVertical interface{} `field:"optional" json:"innerVertical" yaml:"innerVertical"`
	// `CfnTemplate.TableSideBorderOptionsProperty.Left`.
	Left interface{} `field:"optional" json:"left" yaml:"left"`
	// `CfnTemplate.TableSideBorderOptionsProperty.Right`.
	Right interface{} `field:"optional" json:"right" yaml:"right"`
	// `CfnTemplate.TableSideBorderOptionsProperty.Top`.
	Top interface{} `field:"optional" json:"top" yaml:"top"`
}


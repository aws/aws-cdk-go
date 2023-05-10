package awsquicksight


// The side border options for a table.
//
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
	// The table border options of the bottom border.
	Bottom interface{} `field:"optional" json:"bottom" yaml:"bottom"`
	// The table border options of the inner horizontal border.
	InnerHorizontal interface{} `field:"optional" json:"innerHorizontal" yaml:"innerHorizontal"`
	// The table border options of the inner vertical border.
	InnerVertical interface{} `field:"optional" json:"innerVertical" yaml:"innerVertical"`
	// The table border options of the left border.
	Left interface{} `field:"optional" json:"left" yaml:"left"`
	// The table border options of the right border.
	Right interface{} `field:"optional" json:"right" yaml:"right"`
	// The table border options of the top border.
	Top interface{} `field:"optional" json:"top" yaml:"top"`
}


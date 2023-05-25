package awsquicksight


// The border options for a table border.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableBorderOptionsProperty := &TableBorderOptionsProperty{
//   	Color: jsii.String("color"),
//   	Style: jsii.String("style"),
//   	Thickness: jsii.Number(123),
//   }
//
type CfnAnalysis_TableBorderOptionsProperty struct {
	// The color of a table border.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The style (none, solid) of a table border.
	Style *string `field:"optional" json:"style" yaml:"style"`
	// The thickness of a table border.
	Thickness *float64 `field:"optional" json:"thickness" yaml:"thickness"`
}


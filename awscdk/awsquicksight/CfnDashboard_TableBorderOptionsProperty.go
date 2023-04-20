package awsquicksight


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
type CfnDashboard_TableBorderOptionsProperty struct {
	// `CfnDashboard.TableBorderOptionsProperty.Color`.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// `CfnDashboard.TableBorderOptionsProperty.Style`.
	Style *string `field:"optional" json:"style" yaml:"style"`
	// `CfnDashboard.TableBorderOptionsProperty.Thickness`.
	Thickness *float64 `field:"optional" json:"thickness" yaml:"thickness"`
}


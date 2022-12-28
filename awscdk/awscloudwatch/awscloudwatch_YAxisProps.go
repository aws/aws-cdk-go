package awscloudwatch


// Properties for a Y-Axis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   yAxisProps := &yAxisProps{
//   	label: jsii.String("label"),
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   	showUnits: jsii.Boolean(false),
//   }
//
// Experimental.
type YAxisProps struct {
	// The label.
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The max value.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The min value.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// Whether to show units.
	// Experimental.
	ShowUnits *bool `field:"optional" json:"showUnits" yaml:"showUnits"`
}


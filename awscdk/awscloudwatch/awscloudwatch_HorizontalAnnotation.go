package awscloudwatch


// Horizontal annotation to be added to a graph.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   horizontalAnnotation := &horizontalAnnotation{
//   	value: jsii.Number(123),
//
//   	// the properties below are optional
//   	color: jsii.String("color"),
//   	fill: awscdk.Aws_cloudwatch.shading_NONE,
//   	label: jsii.String("label"),
//   	visible: jsii.Boolean(false),
//   }
//
type HorizontalAnnotation struct {
	// The value of the annotation.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to be used for the annotation. The `Color` class has a set of standard colors that can be used here.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Add shading above or below the annotation.
	Fill Shading `field:"optional" json:"fill" yaml:"fill"`
	// Label for the annotation.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Whether the annotation is visible.
	Visible *bool `field:"optional" json:"visible" yaml:"visible"`
}


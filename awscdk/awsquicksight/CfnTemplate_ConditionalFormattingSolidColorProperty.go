package awsquicksight


// Formatting configuration for solid color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingSolidColorProperty := &ConditionalFormattingSolidColorProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   }
//
type CfnTemplate_ConditionalFormattingSolidColorProperty struct {
	// The expression that determines the formatting configuration for solid color.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Determines the color.
	Color *string `field:"optional" json:"color" yaml:"color"`
}


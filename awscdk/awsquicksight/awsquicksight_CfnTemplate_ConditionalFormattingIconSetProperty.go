package awsquicksight


// Formatting configuration for icon set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingIconSetProperty := &ConditionalFormattingIconSetProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	IconSetType: jsii.String("iconSetType"),
//   }
//
type CfnTemplate_ConditionalFormattingIconSetProperty struct {
	// The expression that determines the formatting configuration for the icon set.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Determines the icon set type.
	IconSetType *string `field:"optional" json:"iconSetType" yaml:"iconSetType"`
}


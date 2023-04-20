package awsquicksight


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
	// `CfnTemplate.ConditionalFormattingSolidColorProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnTemplate.ConditionalFormattingSolidColorProperty.Color`.
	Color *string `field:"optional" json:"color" yaml:"color"`
}


package awsquicksight


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
type CfnDashboard_ConditionalFormattingIconSetProperty struct {
	// `CfnDashboard.ConditionalFormattingIconSetProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnDashboard.ConditionalFormattingIconSetProperty.IconSetType`.
	IconSetType *string `field:"optional" json:"iconSetType" yaml:"iconSetType"`
}


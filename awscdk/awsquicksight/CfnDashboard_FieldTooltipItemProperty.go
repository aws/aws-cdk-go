package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldTooltipItemProperty := &FieldTooltipItemProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	Label: jsii.String("label"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_FieldTooltipItemProperty struct {
	// `CfnDashboard.FieldTooltipItemProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnDashboard.FieldTooltipItemProperty.Label`.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// `CfnDashboard.FieldTooltipItemProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}


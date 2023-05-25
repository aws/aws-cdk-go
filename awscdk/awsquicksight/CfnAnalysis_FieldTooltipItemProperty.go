package awsquicksight


// The tooltip item for the fields.
//
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
type CfnAnalysis_FieldTooltipItemProperty struct {
	// The unique ID of the field that is targeted by the tooltip.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The label of the tooltip item.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The visibility of the tooltip item.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}


package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelNameConditionProperty := &labelNameConditionProperty{
//   	labelName: jsii.String("labelName"),
//   }
//
type CfnLoggingConfiguration_LabelNameConditionProperty struct {
	// `CfnLoggingConfiguration.LabelNameConditionProperty.LabelName`.
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
}


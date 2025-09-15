package awssecurityhub


// Enables filtering of security findings based on string field values in OCSF.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocsfStringFilterProperty := &OcsfStringFilterProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Filter: &StringFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfstringfilter.html
//
type CfnAutomationRuleV2_OcsfStringFilterProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfstringfilter.html#cfn-securityhub-automationrulev2-ocsfstringfilter-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Enables filtering of security findings based on string field values in OCSF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfstringfilter.html#cfn-securityhub-automationrulev2-ocsfstringfilter-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}


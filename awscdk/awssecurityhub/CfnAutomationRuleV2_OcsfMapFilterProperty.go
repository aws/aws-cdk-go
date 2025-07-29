package awssecurityhub


// Enables filtering of security findings based on map field values in OCSF.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocsfMapFilterProperty := &OcsfMapFilterProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Filter: &MapFilterProperty{
//   		Comparison: jsii.String("comparison"),
//   		Key: jsii.String("key"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfmapfilter.html
//
type CfnAutomationRuleV2_OcsfMapFilterProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfmapfilter.html#cfn-securityhub-automationrulev2-ocsfmapfilter-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Enables filtering of security findings based on map field values in OCSF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfmapfilter.html#cfn-securityhub-automationrulev2-ocsfmapfilter-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}


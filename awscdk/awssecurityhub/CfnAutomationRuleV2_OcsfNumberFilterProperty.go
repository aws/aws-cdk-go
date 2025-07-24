package awssecurityhub


// Enables filtering of security findings based on numerical field values in OCSF.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ocsfNumberFilterProperty := &OcsfNumberFilterProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Filter: &NumberFilterProperty{
//   		Eq: jsii.Number(123),
//   		Gte: jsii.Number(123),
//   		Lte: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfnumberfilter.html
//
type CfnAutomationRuleV2_OcsfNumberFilterProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfnumberfilter.html#cfn-securityhub-automationrulev2-ocsfnumberfilter-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Enables filtering of security findings based on numerical field values in OCSF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfnumberfilter.html#cfn-securityhub-automationrulev2-ocsfnumberfilter-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}


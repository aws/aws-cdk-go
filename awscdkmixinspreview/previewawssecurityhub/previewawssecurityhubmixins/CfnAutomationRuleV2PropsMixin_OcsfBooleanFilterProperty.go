package previewawssecurityhubmixins


// Enables filtering of security findings based on boolean field values in OCSF.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ocsfBooleanFilterProperty := &OcsfBooleanFilterProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Filter: &BooleanFilterProperty{
//   		Value: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfbooleanfilter.html
//
type CfnAutomationRuleV2PropsMixin_OcsfBooleanFilterProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfbooleanfilter.html#cfn-securityhub-automationrulev2-ocsfbooleanfilter-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Enables filtering of security findings based on boolean field values in OCSF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfbooleanfilter.html#cfn-securityhub-automationrulev2-ocsfbooleanfilter-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}


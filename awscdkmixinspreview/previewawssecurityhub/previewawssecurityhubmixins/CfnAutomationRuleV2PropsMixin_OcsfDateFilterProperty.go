package previewawssecurityhubmixins


// Enables filtering of security findings based on date and timestamp fields in OCSF.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ocsfDateFilterProperty := &OcsfDateFilterProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Filter: &DateFilterProperty{
//   		DateRange: &DateRangeProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		End: jsii.String("end"),
//   		Start: jsii.String("start"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfdatefilter.html
//
type CfnAutomationRuleV2PropsMixin_OcsfDateFilterProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfdatefilter.html#cfn-securityhub-automationrulev2-ocsfdatefilter-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Enables filtering of security findings based on date and timestamp fields in OCSF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-ocsfdatefilter.html#cfn-securityhub-automationrulev2-ocsfdatefilter-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}


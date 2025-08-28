package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProtectProperty := &DataProtectProperty{
//   	Action: jsii.String("action"),
//   	Field: &FieldToProtectProperty{
//   		FieldType: jsii.String("fieldType"),
//
//   		// the properties below are optional
//   		FieldKeys: []*string{
//   			jsii.String("fieldKeys"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ExcludeRateBasedDetails: jsii.Boolean(false),
//   	ExcludeRuleMatchDetails: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html
//
type CfnWebACL_DataProtectProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// Field in log to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-field
	//
	Field interface{} `field:"required" json:"field" yaml:"field"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-excluderatebaseddetails
	//
	ExcludeRateBasedDetails interface{} `field:"optional" json:"excludeRateBasedDetails" yaml:"excludeRateBasedDetails"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-excluderulematchdetails
	//
	ExcludeRuleMatchDetails interface{} `field:"optional" json:"excludeRuleMatchDetails" yaml:"excludeRuleMatchDetails"`
}


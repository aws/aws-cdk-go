package previewawswafv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataProtectProperty := &DataProtectProperty{
//   	Action: jsii.String("action"),
//   	ExcludeRateBasedDetails: jsii.Boolean(false),
//   	ExcludeRuleMatchDetails: jsii.Boolean(false),
//   	Field: &FieldToProtectProperty{
//   		FieldKeys: []*string{
//   			jsii.String("fieldKeys"),
//   		},
//   		FieldType: jsii.String("fieldType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html
//
type CfnWebACLPropsMixin_DataProtectProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-excluderatebaseddetails
	//
	ExcludeRateBasedDetails interface{} `field:"optional" json:"excludeRateBasedDetails" yaml:"excludeRateBasedDetails"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-excluderulematchdetails
	//
	ExcludeRuleMatchDetails interface{} `field:"optional" json:"excludeRuleMatchDetails" yaml:"excludeRuleMatchDetails"`
	// Field in log to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotect.html#cfn-wafv2-webacl-dataprotect-field
	//
	Field interface{} `field:"optional" json:"field" yaml:"field"`
}


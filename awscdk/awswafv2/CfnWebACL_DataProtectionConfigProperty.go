package awswafv2


// Specifies data protection to apply to the web request data for the web ACL.
//
// This is a web ACL level data protection option.
//
// The data protection that you configure for the web ACL alters the data that's available for any other data collection activity, including your AWS WAF logging destinations, web ACL request sampling, and Amazon Security Lake data collection and management. Your other option for data protection is in the logging configuration, which only affects logging.
//
// This is part of the data protection configuration for a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProtectionConfigProperty := &DataProtectionConfigProperty{
//   	DataProtections: []interface{}{
//   		&DataProtectProperty{
//   			Action: jsii.String("action"),
//   			Field: &FieldToProtectProperty{
//   				FieldType: jsii.String("fieldType"),
//
//   				// the properties below are optional
//   				FieldKeys: []*string{
//   					jsii.String("fieldKeys"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ExcludeRateBasedDetails: jsii.Boolean(false),
//   			ExcludeRuleMatchDetails: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotectionconfig.html
//
type CfnWebACL_DataProtectionConfigProperty struct {
	// An array of data protection configurations for specific web request field types.
	//
	// This is defined for each web ACL. AWS WAF applies the specified protection to all web requests that the web ACL inspects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-dataprotectionconfig.html#cfn-wafv2-webacl-dataprotectionconfig-dataprotections
	//
	DataProtections interface{} `field:"required" json:"dataProtections" yaml:"dataProtections"`
}


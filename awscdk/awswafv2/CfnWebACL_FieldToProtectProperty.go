package awswafv2


// Specifies a field type and keys to protect in stored web request data.
//
// This is part of the data protection configuration for a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldToProtectProperty := &FieldToProtectProperty{
//   	FieldType: jsii.String("fieldType"),
//
//   	// the properties below are optional
//   	FieldKeys: []*string{
//   		jsii.String("fieldKeys"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtoprotect.html
//
type CfnWebACL_FieldToProtectProperty struct {
	// Specifies the web request component type to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtoprotect.html#cfn-wafv2-webacl-fieldtoprotect-fieldtype
	//
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
	// Specifies the keys to protect for the specified field type.
	//
	// If you don't specify any key, then all keys for the field type are protected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtoprotect.html#cfn-wafv2-webacl-fieldtoprotect-fieldkeys
	//
	FieldKeys *[]*string `field:"optional" json:"fieldKeys" yaml:"fieldKeys"`
}


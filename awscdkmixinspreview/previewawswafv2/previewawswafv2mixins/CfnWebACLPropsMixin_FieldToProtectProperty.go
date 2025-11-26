package previewawswafv2mixins


// Specifies a field type and keys to protect in stored web request data.
//
// This is part of the data protection configuration for a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldToProtectProperty := &FieldToProtectProperty{
//   	FieldKeys: []*string{
//   		jsii.String("fieldKeys"),
//   	},
//   	FieldType: jsii.String("fieldType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtoprotect.html
//
type CfnWebACLPropsMixin_FieldToProtectProperty struct {
	// Specifies the keys to protect for the specified field type.
	//
	// If you don't specify any key, then all keys for the field type are protected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtoprotect.html#cfn-wafv2-webacl-fieldtoprotect-fieldkeys
	//
	FieldKeys *[]*string `field:"optional" json:"fieldKeys" yaml:"fieldKeys"`
	// Specifies the web request component type to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtoprotect.html#cfn-wafv2-webacl-fieldtoprotect-fieldtype
	//
	FieldType *string `field:"optional" json:"fieldType" yaml:"fieldType"`
}


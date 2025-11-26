package previewawsrbinmixins


// [Tag-level retention rules only] Information about the resource tags used to identify resources that are retained by the retention rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceTagProperty := &ResourceTagProperty{
//   	ResourceTagKey: jsii.String("resourceTagKey"),
//   	ResourceTagValue: jsii.String("resourceTagValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-resourcetag.html
//
type CfnRulePropsMixin_ResourceTagProperty struct {
	// The tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-resourcetag.html#cfn-rbin-rule-resourcetag-resourcetagkey
	//
	ResourceTagKey *string `field:"optional" json:"resourceTagKey" yaml:"resourceTagKey"`
	// The tag value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rbin-rule-resourcetag.html#cfn-rbin-rule-resourcetag-resourcetagvalue
	//
	ResourceTagValue *string `field:"optional" json:"resourceTagValue" yaml:"resourceTagValue"`
}


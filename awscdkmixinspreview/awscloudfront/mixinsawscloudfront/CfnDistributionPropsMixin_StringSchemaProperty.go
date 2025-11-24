package mixinsawscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stringSchemaProperty := &StringSchemaProperty{
//   	Comment: jsii.String("comment"),
//   	DefaultValue: jsii.String("defaultValue"),
//   	Required: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html
//
type CfnDistributionPropsMixin_StringSchemaProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html#cfn-cloudfront-distribution-stringschema-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html#cfn-cloudfront-distribution-stringschema-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html#cfn-cloudfront-distribution-stringschema-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}


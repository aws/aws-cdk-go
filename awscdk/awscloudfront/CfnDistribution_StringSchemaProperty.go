package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringSchemaProperty := &StringSchemaProperty{
//   	Required: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   	DefaultValue: jsii.String("defaultValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html
//
type CfnDistribution_StringSchemaProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html#cfn-cloudfront-distribution-stringschema-required
	//
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html#cfn-cloudfront-distribution-stringschema-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-stringschema.html#cfn-cloudfront-distribution-stringschema-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}


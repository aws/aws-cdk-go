package awswafv2


// Text Transformation applied before parsing the query string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   preParseTextTransformationProperty := &PreParseTextTransformationProperty{
//   	Priority: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-preparsetexttransformation.html
//
type CfnWebACLPropsMixin_PreParseTextTransformationProperty struct {
	// Priority of PreParseTextTransformation being applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-preparsetexttransformation.html#cfn-wafv2-webacl-preparsetexttransformation-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Type of pre-parse text transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-preparsetexttransformation.html#cfn-wafv2-webacl-preparsetexttransformation-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


package awsmediatailor


// For SCTE35_ENHANCED output, defines a key and corresponding value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   keyValuePairProperty := &KeyValuePairProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-keyvaluepair.html
//
type CfnProgramPropsMixin_KeyValuePairProperty struct {
	// The key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-keyvaluepair.html#cfn-mediatailor-program-keyvaluepair-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-keyvaluepair.html#cfn-mediatailor-program-keyvaluepair-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


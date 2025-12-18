package previewawsdevopsagentmixins


// A key-value pair for tags.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyValuePairProperty := &KeyValuePairProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-keyvaluepair.html
//
type CfnAssociationPropsMixin_KeyValuePairProperty struct {
	// The key name of the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-keyvaluepair.html#cfn-devopsagent-association-keyvaluepair-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-keyvaluepair.html#cfn-devopsagent-association-keyvaluepair-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


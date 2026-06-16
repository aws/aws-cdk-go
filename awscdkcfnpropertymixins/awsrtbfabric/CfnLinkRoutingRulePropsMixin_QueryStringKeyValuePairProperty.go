package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   queryStringKeyValuePairProperty := &QueryStringKeyValuePairProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-querystringkeyvaluepair.html
//
type CfnLinkRoutingRulePropsMixin_QueryStringKeyValuePairProperty struct {
	// Query string key — RFC 3986 unreserved characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-querystringkeyvaluepair.html#cfn-rtbfabric-linkroutingrule-querystringkeyvaluepair-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Query string value — RFC 3986 unreserved characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-linkroutingrule-querystringkeyvaluepair.html#cfn-rtbfabric-linkroutingrule-querystringkeyvaluepair-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


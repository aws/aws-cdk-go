package previewawsnetworkfirewallmixins


// Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet.
//
// This setting defines a CloudWatch dimension value to be published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publishMetricActionProperty := &PublishMetricActionProperty{
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-publishmetricaction.html
//
type CfnRuleGroupPropsMixin_PublishMetricActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-publishmetricaction.html#cfn-networkfirewall-rulegroup-publishmetricaction-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}


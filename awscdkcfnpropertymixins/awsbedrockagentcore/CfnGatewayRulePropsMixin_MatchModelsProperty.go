package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   matchModelsProperty := &MatchModelsProperty{
//   	AnyOf: []interface{}{
//   		&MatchModelEntryProperty{
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodels.html
//
type CfnGatewayRulePropsMixin_MatchModelsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodels.html#cfn-bedrockagentcore-gatewayrule-matchmodels-anyof
	//
	AnyOf interface{} `field:"optional" json:"anyOf" yaml:"anyOf"`
}


package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   matchModelEntryProperty := &MatchModelEntryProperty{
//   	Operator: jsii.String("operator"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodelentry.html
//
type CfnGatewayRulePropsMixin_MatchModelEntryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodelentry.html#cfn-bedrockagentcore-gatewayrule-matchmodelentry-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodelentry.html#cfn-bedrockagentcore-gatewayrule-matchmodelentry-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


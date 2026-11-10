package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchModelEntryProperty := &MatchModelEntryProperty{
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodelentry.html
//
type CfnGatewayRule_MatchModelEntryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodelentry.html#cfn-bedrockagentcore-gatewayrule-matchmodelentry-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodelentry.html#cfn-bedrockagentcore-gatewayrule-matchmodelentry-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}


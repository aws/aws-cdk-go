package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchModelsProperty := &MatchModelsProperty{
//   	AnyOf: []interface{}{
//   		&MatchModelEntryProperty{
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			Operator: jsii.String("operator"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodels.html
//
type CfnGatewayRule_MatchModelsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchmodels.html#cfn-bedrockagentcore-gatewayrule-matchmodels-anyof
	//
	AnyOf interface{} `field:"required" json:"anyOf" yaml:"anyOf"`
}


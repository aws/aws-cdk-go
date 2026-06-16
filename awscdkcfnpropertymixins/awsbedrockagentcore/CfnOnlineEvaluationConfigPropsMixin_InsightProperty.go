package awsbedrockagentcore


// An insight configuration for failure analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   insightProperty := &InsightProperty{
//   	InsightId: jsii.String("insightId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-insight.html
//
type CfnOnlineEvaluationConfigPropsMixin_InsightProperty struct {
	// The unique identifier of the insight.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-insight.html#cfn-bedrockagentcore-onlineevaluationconfig-insight-insightid
	//
	InsightId *string `field:"optional" json:"insightId" yaml:"insightId"`
}


package awsbedrock


// Routing criteria for a prompt router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingCriteriaProperty := &RoutingCriteriaProperty{
//   	ResponseQualityDifference: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-defaultpromptrouter-routingcriteria.html
//
type CfnDefaultPromptRouter_RoutingCriteriaProperty struct {
	// The response quality difference threshold for routing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-defaultpromptrouter-routingcriteria.html#cfn-bedrock-defaultpromptrouter-routingcriteria-responsequalitydifference
	//
	ResponseQualityDifference *float64 `field:"required" json:"responseQualityDifference" yaml:"responseQualityDifference"`
}


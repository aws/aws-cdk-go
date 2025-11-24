package mixinsawsbedrock


// Routing criteria for a prompt router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routingCriteriaProperty := &RoutingCriteriaProperty{
//   	ResponseQualityDifference: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-intelligentpromptrouter-routingcriteria.html
//
type CfnIntelligentPromptRouterPropsMixin_RoutingCriteriaProperty struct {
	// The criteria's response quality difference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-intelligentpromptrouter-routingcriteria.html#cfn-bedrock-intelligentpromptrouter-routingcriteria-responsequalitydifference
	//
	ResponseQualityDifference *float64 `field:"optional" json:"responseQualityDifference" yaml:"responseQualityDifference"`
}


package awsbedrock


// Guardrail tier config for topic policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicsTierConfigProperty := &TopicsTierConfigProperty{
//   	TierName: jsii.String("tierName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicstierconfig.html
//
type CfnGuardrail_TopicsTierConfigProperty struct {
	// Tier name for tier configuration in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicstierconfig.html#cfn-bedrock-guardrail-topicstierconfig-tiername
	//
	TierName *string `field:"required" json:"tierName" yaml:"tierName"`
}


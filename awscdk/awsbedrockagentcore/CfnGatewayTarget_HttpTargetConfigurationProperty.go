package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpTargetConfigurationProperty := &HttpTargetConfigurationProperty{
//   	AgentcoreRuntime: &RuntimeTargetConfigurationProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		Qualifier: jsii.String("qualifier"),
//   		Schema: &HttpApiSchemaConfigurationProperty{
//   			Source: &ApiSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   		},
//   	},
//   	Passthrough: &PassthroughTargetConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		ProtocolType: jsii.String("protocolType"),
//
//   		// the properties below are optional
//   		Schema: &HttpApiSchemaConfigurationProperty{
//   			Source: &ApiSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   		},
//   		StickinessConfiguration: &StickinessConfigurationProperty{
//   			Identifier: jsii.String("identifier"),
//
//   			// the properties below are optional
//   			Timeout: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httptargetconfiguration.html
//
type CfnGatewayTarget_HttpTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-httptargetconfiguration-agentcoreruntime
	//
	AgentcoreRuntime interface{} `field:"optional" json:"agentcoreRuntime" yaml:"agentcoreRuntime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-httptargetconfiguration-passthrough
	//
	Passthrough interface{} `field:"optional" json:"passthrough" yaml:"passthrough"`
}


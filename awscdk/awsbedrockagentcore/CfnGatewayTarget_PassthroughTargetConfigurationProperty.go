package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   passthroughTargetConfigurationProperty := &PassthroughTargetConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	ProtocolType: jsii.String("protocolType"),
//
//   	// the properties below are optional
//   	Schema: &HttpApiSchemaConfigurationProperty{
//   		Source: &ApiSchemaConfigurationProperty{
//   			InlinePayload: jsii.String("inlinePayload"),
//   			S3: &S3ConfigurationProperty{
//   				BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   	},
//   	StickinessConfiguration: &StickinessConfigurationProperty{
//   		Identifier: jsii.String("identifier"),
//
//   		// the properties below are optional
//   		Timeout: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration.html
//
type CfnGatewayTarget_PassthroughTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration-protocoltype
	//
	ProtocolType *string `field:"required" json:"protocolType" yaml:"protocolType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-passthroughtargetconfiguration-stickinessconfiguration
	//
	StickinessConfiguration interface{} `field:"optional" json:"stickinessConfiguration" yaml:"stickinessConfiguration"`
}


package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mcpToolSchemaConfigurationProperty := &McpToolSchemaConfigurationProperty{
//   	InlinePayload: jsii.String("inlinePayload"),
//   	S3: &S3ConfigurationProperty{
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptoolschemaconfiguration.html
//
type CfnGatewayTarget_McpToolSchemaConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptoolschemaconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcptoolschemaconfiguration-inlinepayload
	//
	InlinePayload *string `field:"optional" json:"inlinePayload" yaml:"inlinePayload"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptoolschemaconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcptoolschemaconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}


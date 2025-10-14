package awsbedrockagentcore


// The API schema configuration for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiSchemaConfigurationProperty := &ApiSchemaConfigurationProperty{
//   	InlinePayload: jsii.String("inlinePayload"),
//   	S3: &S3ConfigurationProperty{
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apischemaconfiguration.html
//
type CfnGatewayTarget_ApiSchemaConfigurationProperty struct {
	// The inline payload for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apischemaconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apischemaconfiguration-inlinepayload
	//
	InlinePayload *string `field:"optional" json:"inlinePayload" yaml:"inlinePayload"`
	// The API schema configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-apischemaconfiguration.html#cfn-bedrockagentcore-gatewaytarget-apischemaconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}


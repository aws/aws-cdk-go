package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiSchemaConfigurationProperty := &HttpApiSchemaConfigurationProperty{
//   	Source: &ApiSchemaConfigurationProperty{
//   		InlinePayload: jsii.String("inlinePayload"),
//   		S3: &S3ConfigurationProperty{
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpapischemaconfiguration.html
//
type CfnGatewayTarget_HttpApiSchemaConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httpapischemaconfiguration.html#cfn-bedrockagentcore-gatewaytarget-httpapischemaconfiguration-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
}


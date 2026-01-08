package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	AllowedQueryParameters: []*string{
//   		jsii.String("allowedQueryParameters"),
//   	},
//   	AllowedRequestHeaders: []*string{
//   		jsii.String("allowedRequestHeaders"),
//   	},
//   	AllowedResponseHeaders: []*string{
//   		jsii.String("allowedResponseHeaders"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-metadataconfiguration.html
//
type CfnGatewayTarget_MetadataConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-metadataconfiguration.html#cfn-bedrockagentcore-gatewaytarget-metadataconfiguration-allowedqueryparameters
	//
	AllowedQueryParameters *[]*string `field:"optional" json:"allowedQueryParameters" yaml:"allowedQueryParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-metadataconfiguration.html#cfn-bedrockagentcore-gatewaytarget-metadataconfiguration-allowedrequestheaders
	//
	AllowedRequestHeaders *[]*string `field:"optional" json:"allowedRequestHeaders" yaml:"allowedRequestHeaders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-metadataconfiguration.html#cfn-bedrockagentcore-gatewaytarget-metadataconfiguration-allowedresponseheaders
	//
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
}


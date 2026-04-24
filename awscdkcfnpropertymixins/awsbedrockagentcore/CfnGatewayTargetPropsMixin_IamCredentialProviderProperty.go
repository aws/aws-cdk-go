package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   iamCredentialProviderProperty := &IamCredentialProviderProperty{
//   	Region: jsii.String("region"),
//   	Service: jsii.String("service"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.html
//
type CfnGatewayTargetPropsMixin_IamCredentialProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-iamcredentialprovider-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-iamcredentialprovider-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}


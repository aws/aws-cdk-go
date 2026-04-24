package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamCredentialProviderProperty := &IamCredentialProviderProperty{
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.html
//
type CfnGatewayTarget_IamCredentialProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-iamcredentialprovider-service
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.html#cfn-bedrockagentcore-gatewaytarget-iamcredentialprovider-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}


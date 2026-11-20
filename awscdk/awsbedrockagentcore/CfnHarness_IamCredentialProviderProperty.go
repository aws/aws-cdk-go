package awsbedrockagentcore


// IAM credentials used to SigV4-sign requests.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-iamcredentialprovider.html
//
type CfnHarness_IamCredentialProviderProperty struct {
	// The AWS service name used for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-iamcredentialprovider.html#cfn-bedrockagentcore-harness-iamcredentialprovider-service
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// The AWS Region used for SigV4 signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-iamcredentialprovider.html#cfn-bedrockagentcore-harness-iamcredentialprovider-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}


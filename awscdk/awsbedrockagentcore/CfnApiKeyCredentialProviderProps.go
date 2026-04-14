package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApiKeyCredentialProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiKeyCredentialProviderProps := &CfnApiKeyCredentialProviderProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ApiKey: jsii.String("apiKey"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html
//
type CfnApiKeyCredentialProviderProps struct {
	// The name of the API key credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API key to use for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Tags to assign to the API key credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-apikeycredentialprovider.html#cfn-bedrockagentcore-apikeycredentialprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


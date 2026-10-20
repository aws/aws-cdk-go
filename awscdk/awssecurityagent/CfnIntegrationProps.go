package awssecurityagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationProps := &CfnIntegrationProps{
//   	DisplayName: jsii.String("displayName"),
//   	Provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PrivateConnectionName: jsii.String("privateConnectionName"),
//   	ProviderInput: &ProviderInputProperty{
//   		Bitbucket: &BitbucketInputProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   			State: jsii.String("state"),
//   			Workspace: jsii.String("workspace"),
//   		},
//   		Confluence: &ConfluenceInputProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   			SiteUrl: jsii.String("siteUrl"),
//   			State: jsii.String("state"),
//   		},
//   		Github: &GitHubInputProperty{
//   			Code: jsii.String("code"),
//   			State: jsii.String("state"),
//
//   			// the properties below are optional
//   			InstallationId: jsii.String("installationId"),
//   			OrganizationName: jsii.String("organizationName"),
//   			TargetUrl: jsii.String("targetUrl"),
//   		},
//   		Gitlab: &GitLabInputProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			TokenType: jsii.String("tokenType"),
//
//   			// the properties below are optional
//   			GroupId: jsii.String("groupId"),
//   			TargetUrl: jsii.String("targetUrl"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html
//
type CfnIntegrationProps struct {
	// The display name for the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html#cfn-securityagent-integration-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The integration provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html#cfn-securityagent-integration-provider
	//
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The identifier of the AWS KMS key to use for encrypting data associated with the integration.
	//
	// Can be a key ID, key ARN, alias name, or alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html#cfn-securityagent-integration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of an active private connection used to reach a self-hosted provider instance over private networking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html#cfn-securityagent-integration-privateconnectionname
	//
	PrivateConnectionName *string `field:"optional" json:"privateConnectionName" yaml:"privateConnectionName"`
	// The provider-specific input required to create the integration.
	//
	// Exactly one of the provider input objects must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html#cfn-securityagent-integration-providerinput
	//
	ProviderInput interface{} `field:"optional" json:"providerInput" yaml:"providerInput"`
	// The tags to associate with the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-integration.html#cfn-securityagent-integration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


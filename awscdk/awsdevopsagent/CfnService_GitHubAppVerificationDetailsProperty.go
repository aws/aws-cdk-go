package awsdevopsagent


// The token verifies an existing DevOps Agent GitHub App installation on github.com and is then discarded. Not supported for GitHub Enterprise Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubAppVerificationDetailsProperty := &GitHubAppVerificationDetailsProperty{
//   	AppVariant: jsii.String("appVariant"),
//   	TokenValue: jsii.String("tokenValue"),
//
//   	// the properties below are optional
//   	OrganizationName: jsii.String("organizationName"),
//   	TargetUrl: jsii.String("targetUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubappverificationdetails.html
//
type CfnService_GitHubAppVerificationDetailsProperty struct {
	// Permission variant of the DevOps Agent GitHub App that must already be installed on the organization or account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubappverificationdetails.html#cfn-devopsagent-service-githubappverificationdetails-appvariant
	//
	AppVariant *string `field:"required" json:"appVariant" yaml:"appVariant"`
	// GitHub personal access token (PAT) value.
	//
	// Accepts classic PATs (ghp_...), fine-grained PATs (github_pat_...), and GitHub Enterprise Server token formats
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubappverificationdetails.html#cfn-devopsagent-service-githubappverificationdetails-tokenvalue
	//
	TokenValue *string `field:"required" json:"tokenValue" yaml:"tokenValue"`
	// GitHub organization to scope the service to.
	//
	// The token must be an active member of the organization and only that organization's repositories are accessible. Omit to register the token's own user account
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubappverificationdetails.html#cfn-devopsagent-service-githubappverificationdetails-organizationname
	//
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// GitHub instance URL.
	//
	// Omit for github.com. Provide the instance host for GitHub Enterprise Server or GitHub Enterprise Cloud with data residency (e.g. https://github.example.com)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubappverificationdetails.html#cfn-devopsagent-service-githubappverificationdetails-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}


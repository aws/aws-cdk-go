package awsdevopsagent


// The token is stored and used as the long-lived credential for GitHub API access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gitHubPatDetailsProperty := &GitHubPatDetailsProperty{
//   	OrganizationName: jsii.String("organizationName"),
//   	TargetUrl: jsii.String("targetUrl"),
//   	TokenValue: jsii.String("tokenValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubpatdetails.html
//
type CfnServicePropsMixin_GitHubPatDetailsProperty struct {
	// GitHub organization to scope the service to.
	//
	// The token must be an active member of the organization and only that organization's repositories are accessible. Omit to register the token's own user account
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubpatdetails.html#cfn-devopsagent-service-githubpatdetails-organizationname
	//
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// GitHub instance URL.
	//
	// Omit for github.com. Provide the instance host for GitHub Enterprise Server or GitHub Enterprise Cloud with data residency (e.g. https://github.example.com)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubpatdetails.html#cfn-devopsagent-service-githubpatdetails-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// GitHub personal access token (PAT) value.
	//
	// Accepts classic PATs (ghp_...), fine-grained PATs (github_pat_...), and GitHub Enterprise Server token formats
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubpatdetails.html#cfn-devopsagent-service-githubpatdetails-tokenvalue
	//
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}


package awssecurityagent


// The input required to create a GitHub integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubInputProperty := &GitHubInputProperty{
//   	Code: jsii.String("code"),
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	InstallationId: jsii.String("installationId"),
//   	OrganizationName: jsii.String("organizationName"),
//   	TargetUrl: jsii.String("targetUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-githubinput.html
//
type CfnIntegration_GitHubInputProperty struct {
	// The OAuth authorization code received from GitHub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-githubinput.html#cfn-securityagent-integration-githubinput-code
	//
	Code *string `field:"required" json:"code" yaml:"code"`
	// The CSRF state token for validating the OAuth flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-githubinput.html#cfn-securityagent-integration-githubinput-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// The installation identifier provided by GitHub Enterprise Server on the install callback.
	//
	// Required for GitHub Enterprise Server integrations and ignored for GitHub.com.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-githubinput.html#cfn-securityagent-integration-githubinput-installationid
	//
	InstallationId *string `field:"optional" json:"installationId" yaml:"installationId"`
	// The name of the GitHub organization to integrate with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-githubinput.html#cfn-securityagent-integration-githubinput-organizationname
	//
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// The HTTPS URL of a self-hosted GitHub Enterprise Server instance.
	//
	// Omit this value for GitHub.com.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-githubinput.html#cfn-securityagent-integration-githubinput-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}


package awssecurityagent


// The input required to create a GitLab integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gitLabInputProperty := &GitLabInputProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	GroupId: jsii.String("groupId"),
//   	TargetUrl: jsii.String("targetUrl"),
//   	TokenType: jsii.String("tokenType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-gitlabinput.html
//
type CfnIntegrationPropsMixin_GitLabInputProperty struct {
	// The GitLab access token used to authenticate.
	//
	// This can be a personal access token or a group access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-gitlabinput.html#cfn-securityagent-integration-gitlabinput-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The identifier of the GitLab group.
	//
	// Required when TokenType is GROUP and ignored for personal tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-gitlabinput.html#cfn-securityagent-integration-gitlabinput-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The HTTPS URL of a self-managed GitLab instance.
	//
	// Omit this value for GitLab SaaS (gitlab.com).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-gitlabinput.html#cfn-securityagent-integration-gitlabinput-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// The type of GitLab access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-gitlabinput.html#cfn-securityagent-integration-gitlabinput-tokentype
	//
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}


package previewawsdevopsagentmixins


// GitLab service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gitLabDetailsProperty := &GitLabDetailsProperty{
//   	GroupId: jsii.String("groupId"),
//   	TargetUrl: jsii.String("targetUrl"),
//   	TokenType: jsii.String("tokenType"),
//   	TokenValue: jsii.String("tokenValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html
//
type CfnServicePropsMixin_GitLabDetailsProperty struct {
	// Optional GitLab group ID for group-level access tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// GitLab instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// Type of GitLab access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-tokentype
	//
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// GitLab access token value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-tokenvalue
	//
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}


package awsdevopsagent


// GitLab service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitLabDetailsProperty := &GitLabDetailsProperty{
//   	TargetUrl: jsii.String("targetUrl"),
//   	TokenType: jsii.String("tokenType"),
//   	TokenValue: jsii.String("tokenValue"),
//
//   	// the properties below are optional
//   	GroupId: jsii.String("groupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html
//
type CfnService_GitLabDetailsProperty struct {
	// GitLab instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-targeturl
	//
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Type of GitLab access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-tokentype
	//
	TokenType *string `field:"required" json:"tokenType" yaml:"tokenType"`
	// GitLab access token value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-tokenvalue
	//
	TokenValue *string `field:"required" json:"tokenValue" yaml:"tokenValue"`
	// Optional GitLab group ID for group-level access tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-gitlabdetails.html#cfn-devopsagent-service-gitlabdetails-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
}


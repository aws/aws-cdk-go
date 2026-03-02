package awsdevopsagent


// GitLab service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredGitLabServiceDetailsProperty := &RegisteredGitLabServiceDetailsProperty{
//   	TargetUrl: jsii.String("targetUrl"),
//   	TokenType: jsii.String("tokenType"),
//
//   	// the properties below are optional
//   	GroupId: jsii.String("groupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html
//
type CfnService_RegisteredGitLabServiceDetailsProperty struct {
	// GitLab instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html#cfn-devopsagent-service-registeredgitlabservicedetails-targeturl
	//
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Type of GitLab access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html#cfn-devopsagent-service-registeredgitlabservicedetails-tokentype
	//
	TokenType *string `field:"required" json:"tokenType" yaml:"tokenType"`
	// Optional GitLab group ID for group-level access tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html#cfn-devopsagent-service-registeredgitlabservicedetails-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
}


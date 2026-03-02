package previewawsdevopsagentmixins


// GitLab service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registeredGitLabServiceDetailsProperty := &RegisteredGitLabServiceDetailsProperty{
//   	GroupId: jsii.String("groupId"),
//   	TargetUrl: jsii.String("targetUrl"),
//   	TokenType: jsii.String("tokenType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html
//
type CfnServicePropsMixin_RegisteredGitLabServiceDetailsProperty struct {
	// Optional GitLab group ID for group-level access tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html#cfn-devopsagent-service-registeredgitlabservicedetails-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// GitLab instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html#cfn-devopsagent-service-registeredgitlabservicedetails-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// Type of GitLab access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgitlabservicedetails.html#cfn-devopsagent-service-registeredgitlabservicedetails-tokentype
	//
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}


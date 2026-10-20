package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubDetailsProperty := &GitHubDetailsProperty{
//   	AppVerification: &GitHubAppVerificationDetailsProperty{
//   		AppVariant: jsii.String("appVariant"),
//   		TokenValue: jsii.String("tokenValue"),
//
//   		// the properties below are optional
//   		OrganizationName: jsii.String("organizationName"),
//   		TargetUrl: jsii.String("targetUrl"),
//   	},
//   	DirectAccess: &GitHubPatDetailsProperty{
//   		TokenValue: jsii.String("tokenValue"),
//
//   		// the properties below are optional
//   		OrganizationName: jsii.String("organizationName"),
//   		TargetUrl: jsii.String("targetUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubdetails.html
//
type CfnService_GitHubDetailsProperty struct {
	// The token verifies an existing DevOps Agent GitHub App installation on github.com and is then discarded. Not supported for GitHub Enterprise Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubdetails.html#cfn-devopsagent-service-githubdetails-appverification
	//
	AppVerification interface{} `field:"optional" json:"appVerification" yaml:"appVerification"`
	// The token is stored and used as the long-lived credential for GitHub API access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-githubdetails.html#cfn-devopsagent-service-githubdetails-directaccess
	//
	DirectAccess interface{} `field:"optional" json:"directAccess" yaml:"directAccess"`
}


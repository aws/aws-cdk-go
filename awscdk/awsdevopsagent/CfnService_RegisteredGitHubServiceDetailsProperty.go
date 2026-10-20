package awsdevopsagent


// GitHub service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredGitHubServiceDetailsProperty := &RegisteredGitHubServiceDetailsProperty{
//   	Owner: jsii.String("owner"),
//   	OwnerType: jsii.String("ownerType"),
//
//   	// the properties below are optional
//   	AppVariant: jsii.String("appVariant"),
//   	AuthMode: jsii.String("authMode"),
//   	HostingType: jsii.String("hostingType"),
//   	TargetUrl: jsii.String("targetUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html
//
type CfnService_RegisteredGitHubServiceDetailsProperty struct {
	// The GitHub repository owner name derived from the token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html#cfn-devopsagent-service-registeredgithubservicedetails-owner
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Type of repository owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html#cfn-devopsagent-service-registeredgithubservicedetails-ownertype
	//
	OwnerType *string `field:"required" json:"ownerType" yaml:"ownerType"`
	// The GitHub App variant the service was registered with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html#cfn-devopsagent-service-registeredgithubservicedetails-appvariant
	//
	AppVariant *string `field:"optional" json:"appVariant" yaml:"appVariant"`
	// How the service authenticates against the GitHub API: app after AppVerification, pat after DirectAccess.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html#cfn-devopsagent-service-registeredgithubservicedetails-authmode
	//
	AuthMode *string `field:"optional" json:"authMode" yaml:"authMode"`
	// The GitHub hosting type the service was registered with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html#cfn-devopsagent-service-registeredgithubservicedetails-hostingtype
	//
	HostingType *string `field:"optional" json:"hostingType" yaml:"hostingType"`
	// GitHub instance URL (absent for github.com).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredgithubservicedetails.html#cfn-devopsagent-service-registeredgithubservicedetails-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}


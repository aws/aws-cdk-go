package awscodedeploy


// `GitHubLocation` is a property of the [CodeDeploy DeploymentGroup Revision](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html) property that specifies the location of an application revision that is stored in GitHub.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubLocationProperty := &GitHubLocationProperty{
//   	CommitId: jsii.String("commitId"),
//   	Repository: jsii.String("repository"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-githublocation.html
//
type CfnDeploymentGroup_GitHubLocationProperty struct {
	// The SHA1 commit ID of the GitHub commit that represents the bundled artifacts for the application revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-githublocation.html#cfn-codedeploy-deploymentgroup-githublocation-commitid
	//
	CommitId *string `field:"required" json:"commitId" yaml:"commitId"`
	// The GitHub account and repository pair that stores a reference to the commit that represents the bundled artifacts for the application revision.
	//
	// Specify the value as `account/repository` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-githublocation.html#cfn-codedeploy-deploymentgroup-githublocation-repository
	//
	Repository *string `field:"required" json:"repository" yaml:"repository"`
}


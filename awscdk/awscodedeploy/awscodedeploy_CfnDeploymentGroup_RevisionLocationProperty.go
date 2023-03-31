package awscodedeploy


// `RevisionLocation` is a property that defines the location of the CodeDeploy application revision to deploy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   revisionLocationProperty := &revisionLocationProperty{
//   	gitHubLocation: &gitHubLocationProperty{
//   		commitId: jsii.String("commitId"),
//   		repository: jsii.String("repository"),
//   	},
//   	revisionType: jsii.String("revisionType"),
//   	s3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		bundleType: jsii.String("bundleType"),
//   		eTag: jsii.String("eTag"),
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnDeploymentGroup_RevisionLocationProperty struct {
	// Information about the location of application artifacts stored in GitHub.
	GitHubLocation interface{} `field:"optional" json:"gitHubLocation" yaml:"gitHubLocation"`
	// The type of application revision:.
	//
	// - S3: An application revision stored in Amazon S3.
	// - GitHub: An application revision stored in GitHub (EC2/On-premises deployments only).
	// - String: A YAML-formatted or JSON-formatted string ( AWS Lambda deployments only).
	// - AppSpecContent: An `AppSpecContent` object that contains the contents of an AppSpec file for an AWS Lambda or Amazon ECS deployment. The content is formatted as JSON or YAML stored as a RawString.
	RevisionType *string `field:"optional" json:"revisionType" yaml:"revisionType"`
	// Information about the location of a revision stored in Amazon S3.
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}


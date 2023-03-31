package awscodedeploy


// `Deployment` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource that specifies an AWS CodeDeploy application revision to be deployed to instances in the deployment group. If you specify an application revision, your target revision is deployed as soon as the provisioning process is complete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentProperty := &deploymentProperty{
//   	revision: &revisionLocationProperty{
//   		gitHubLocation: &gitHubLocationProperty{
//   			commitId: jsii.String("commitId"),
//   			repository: jsii.String("repository"),
//   		},
//   		revisionType: jsii.String("revisionType"),
//   		s3Location: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			bundleType: jsii.String("bundleType"),
//   			eTag: jsii.String("eTag"),
//   			version: jsii.String("version"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	ignoreApplicationStopFailures: jsii.Boolean(false),
//   }
//
type CfnDeploymentGroup_DeploymentProperty struct {
	// Information about the location of stored application artifacts and the service from which to retrieve them.
	Revision interface{} `field:"required" json:"revision" yaml:"revision"`
	// A comment about the deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If true, then if an `ApplicationStop` , `BeforeBlockTraffic` , or `AfterBlockTraffic` deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event.
	//
	// For example, if `ApplicationStop` fails, the deployment continues with DownloadBundle. If `BeforeBlockTraffic` fails, the deployment continues with `BlockTraffic` . If `AfterBlockTraffic` fails, the deployment continues with `ApplicationStop` .
	//
	// If false or not specified, then if a lifecycle event fails during a deployment to an instance, that deployment fails. If deployment to that instance is part of an overall deployment and the number of healthy hosts is not less than the minimum number of healthy hosts, then a deployment to the next instance is attempted.
	//
	// During a deployment, the AWS CodeDeploy agent runs the scripts specified for `ApplicationStop` , `BeforeBlockTraffic` , and `AfterBlockTraffic` in the AppSpec file from the previous successful deployment. (All other scripts are run from the AppSpec file in the current deployment.) If one of these scripts contains an error and does not run successfully, the deployment can fail.
	//
	// If the cause of the failure is a script from the last successful deployment that will never run successfully, create a new deployment and use `ignoreApplicationStopFailures` to specify that the `ApplicationStop` , `BeforeBlockTraffic` , and `AfterBlockTraffic` failures should be ignored.
	IgnoreApplicationStopFailures interface{} `field:"optional" json:"ignoreApplicationStopFailures" yaml:"ignoreApplicationStopFailures"`
}


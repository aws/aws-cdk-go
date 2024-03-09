package awsbatch


// The platform configuration for jobs that are running on Fargate resources.
//
// Jobs that run on Amazon EC2 resources must not specify this parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fargatePlatformConfigurationProperty := &FargatePlatformConfigurationProperty{
//   	PlatformVersion: jsii.String("platformVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-fargateplatformconfiguration.html
//
type CfnJobDefinition_FargatePlatformConfigurationProperty struct {
	// The AWS Fargate platform version where the jobs are running.
	//
	// A platform version is specified only for jobs that are running on Fargate resources. If one isn't specified, the `LATEST` platform version is used by default. This uses a recent, approved version of the AWS Fargate platform for compute resources. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-fargateplatformconfiguration.html#cfn-batch-jobdefinition-fargateplatformconfiguration-platformversion
	//
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
}


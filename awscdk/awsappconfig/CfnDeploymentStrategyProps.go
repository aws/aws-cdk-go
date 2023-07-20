package awsappconfig


// Properties for defining a `CfnDeploymentStrategy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentStrategyProps := &CfnDeploymentStrategyProps{
//   	DeploymentDurationInMinutes: jsii.Number(123),
//   	GrowthFactor: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ReplicateTo: jsii.String("replicateTo"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FinalBakeTimeInMinutes: jsii.Number(123),
//   	GrowthType: jsii.String("growthType"),
//   	Tags: []tagsProperty{
//   		&tagsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html
//
type CfnDeploymentStrategyProps struct {
	// Total amount of time for a deployment to last.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-deploymentdurationinminutes
	//
	DeploymentDurationInMinutes *float64 `field:"required" json:"deploymentDurationInMinutes" yaml:"deploymentDurationInMinutes"`
	// The percentage of targets to receive a deployed configuration during each interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-growthfactor
	//
	GrowthFactor *float64 `field:"required" json:"growthFactor" yaml:"growthFactor"`
	// A name for the deployment strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Save the deployment strategy to a Systems Manager (SSM) document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-replicateto
	//
	ReplicateTo *string `field:"required" json:"replicateTo" yaml:"replicateTo"`
	// A description of the deployment strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the amount of time AWS AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete.
	//
	// If an alarm is triggered during this time, AWS AppConfig rolls back the deployment. You must configure permissions for AWS AppConfig to roll back based on CloudWatch alarms. For more information, see [Configuring permissions for rollback based on Amazon CloudWatch alarms](https://docs.aws.amazon.com/appconfig/latest/userguide/getting-started-with-appconfig-cloudwatch-alarms-permissions.html) in the *AWS AppConfig User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-finalbaketimeinminutes
	//
	FinalBakeTimeInMinutes *float64 `field:"optional" json:"finalBakeTimeInMinutes" yaml:"finalBakeTimeInMinutes"`
	// The algorithm used to define how percentage grows over time. AWS AppConfig supports the following growth types:.
	//
	// *Linear* : For this type, AWS AppConfig processes the deployment by dividing the total number of targets by the value specified for `Step percentage` . For example, a linear deployment that uses a `Step percentage` of 10 deploys the configuration to 10 percent of the hosts. After those deployments are complete, the system deploys the configuration to the next 10 percent. This continues until 100% of the targets have successfully received the configuration.
	//
	// *Exponential* : For this type, AWS AppConfig processes the deployment exponentially using the following formula: `G*(2^N)` . In this formula, `G` is the growth factor specified by the user and `N` is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:
	//
	// `2*(2^0)`
	//
	// `2*(2^1)`
	//
	// `2*(2^2)`
	//
	// Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-growthtype
	//
	GrowthType *string `field:"optional" json:"growthType" yaml:"growthType"`
	// Assigns metadata to an AWS AppConfig resource.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. You can specify a maximum of 50 tags for a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deploymentstrategy.html#cfn-appconfig-deploymentstrategy-tags
	//
	Tags *[]*CfnDeploymentStrategy_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}


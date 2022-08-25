package awsappconfig


// Properties for defining a `CfnDeploymentStrategy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentStrategyProps := &cfnDeploymentStrategyProps{
//   	deploymentDurationInMinutes: jsii.Number(123),
//   	growthFactor: jsii.Number(123),
//   	name: jsii.String("name"),
//   	replicateTo: jsii.String("replicateTo"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	finalBakeTimeInMinutes: jsii.Number(123),
//   	growthType: jsii.String("growthType"),
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeploymentStrategyProps struct {
	// Total amount of time for a deployment to last.
	DeploymentDurationInMinutes *float64 `field:"required" json:"deploymentDurationInMinutes" yaml:"deploymentDurationInMinutes"`
	// The percentage of targets to receive a deployed configuration during each interval.
	GrowthFactor *float64 `field:"required" json:"growthFactor" yaml:"growthFactor"`
	// A name for the deployment strategy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo *string `field:"required" json:"replicateTo" yaml:"replicateTo"`
	// A description of the deployment strategy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The amount of time AWS AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic roll back.
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
	GrowthType *string `field:"optional" json:"growthType" yaml:"growthType"`
	// Assigns metadata to an AWS AppConfig resource.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. You can specify a maximum of 50 tags for a resource.
	Tags *[]*CfnDeploymentStrategy_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}


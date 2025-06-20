package awscodedeploy


// Properties for defining a `CfnDeploymentConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentConfigProps := &CfnDeploymentConfigProps{
//   	ComputePlatform: jsii.String("computePlatform"),
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   	MinimumHealthyHosts: &MinimumHealthyHostsProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	TrafficRoutingConfig: &TrafficRoutingConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		TimeBasedCanary: &TimeBasedCanaryProperty{
//   			CanaryInterval: jsii.Number(123),
//   			CanaryPercentage: jsii.Number(123),
//   		},
//   		TimeBasedLinear: &TimeBasedLinearProperty{
//   			LinearInterval: jsii.Number(123),
//   			LinearPercentage: jsii.Number(123),
//   		},
//   	},
//   	ZonalConfig: &ZonalConfigProperty{
//   		FirstZoneMonitorDurationInSeconds: jsii.Number(123),
//   		MinimumHealthyHostsPerZone: &MinimumHealthyHostsPerZoneProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		MonitorDurationInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
//
type CfnDeploymentConfigProps struct {
	// The destination platform type for the deployment ( `Lambda` , `Server` , or `ECS` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-computeplatform
	//
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A name for the deployment configuration.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-deploymentconfigname
	//
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The minimum number of healthy instances that should be available at any time during the deployment.
	//
	// There are two parameters expected in the input: type and value.
	//
	// The type parameter takes either of the following values:
	//
	// - HOST_COUNT: The value parameter represents the minimum number of healthy instances as an absolute value.
	// - FLEET_PERCENT: The value parameter represents the minimum number of healthy instances as a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
	//
	// The value parameter takes an integer.
	//
	// For example, to set a minimum of 95% healthy instance, specify a type of FLEET_PERCENT and a value of 95.
	//
	// For more information about instance health, see [CodeDeploy Instance Health](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-health.html) in the AWS CodeDeploy User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts
	//
	MinimumHealthyHosts interface{} `field:"optional" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The configuration that specifies how the deployment traffic is routed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-trafficroutingconfig
	//
	TrafficRoutingConfig interface{} `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
	// Configure the `ZonalConfig` object if you want AWS CodeDeploy to deploy your application to one [Availability Zone](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-availability-zones) at a time, within an AWS Region.
	//
	// For more information about the zonal configuration feature, see [zonal configuration](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations-create.html#zonal-config) in the *CodeDeploy User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-zonalconfig
	//
	ZonalConfig interface{} `field:"optional" json:"zonalConfig" yaml:"zonalConfig"`
}


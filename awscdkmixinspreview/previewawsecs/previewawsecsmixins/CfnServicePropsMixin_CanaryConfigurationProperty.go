package previewawsecsmixins


// Configuration for a canary deployment strategy that shifts a fixed percentage of traffic to the new service revision, waits for a specified bake time, then shifts the remaining traffic.
//
// The following validation applies only to Canary deployments created through CloudFormation . CloudFormation operations time out after 36 hours. Canary deployments can approach this limit because of their extended duration. This can cause CloudFormation to roll back the deployment. To prevent timeout-related rollbacks, CloudFormation rejects deployments when the calculated deployment time exceeds 33 hours based on your template configuration:
//
// `BakeTimeInMinutes + CanaryBakeTimeInMinutes`
//
// Additional backend processes (such as task scaling and running lifecycle hooks) can extend deployment time beyond these calculations. Even deployments under the 33-hour threshold might still time out if these processes cause the total duration to exceed 36 hours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   canaryConfigurationProperty := &CanaryConfigurationProperty{
//   	CanaryBakeTimeInMinutes: jsii.Number(123),
//   	CanaryPercent: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-canaryconfiguration.html
//
type CfnServicePropsMixin_CanaryConfigurationProperty struct {
	// The amount of time in minutes to wait during the canary phase before shifting the remaining production traffic to the new service revision.
	//
	// Valid values are 0 to 1440 minutes (24 hours). The default value is 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-canaryconfiguration.html#cfn-ecs-service-canaryconfiguration-canarybaketimeinminutes
	//
	CanaryBakeTimeInMinutes *float64 `field:"optional" json:"canaryBakeTimeInMinutes" yaml:"canaryBakeTimeInMinutes"`
	// The percentage of production traffic to shift to the new service revision during the canary phase.
	//
	// Valid values are multiples of 0.1 from 0.1 to 100.0. The default value is 5.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-canaryconfiguration.html#cfn-ecs-service-canaryconfiguration-canarypercent
	//
	CanaryPercent *float64 `field:"optional" json:"canaryPercent" yaml:"canaryPercent"`
}


package awscdk


// Construction properties of `CfnCodeDeployBlueGreenHook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenHookProps := &CfnCodeDeployBlueGreenHookProps{
//   	Applications: []cfnCodeDeployBlueGreenApplication{
//   		&cfnCodeDeployBlueGreenApplication{
//   			EcsAttributes: &CfnCodeDeployBlueGreenEcsAttributes{
//   				TaskDefinitions: []*string{
//   					jsii.String("taskDefinitions"),
//   				},
//   				TaskSets: []*string{
//   					jsii.String("taskSets"),
//   				},
//   				TrafficRouting: &CfnTrafficRouting{
//   					ProdTrafficRoute: &CfnTrafficRoute{
//   						LogicalId: jsii.String("logicalId"),
//   						Type: jsii.String("type"),
//   					},
//   					TargetGroups: []*string{
//   						jsii.String("targetGroups"),
//   					},
//   					TestTrafficRoute: &CfnTrafficRoute{
//   						LogicalId: jsii.String("logicalId"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			Target: &CfnCodeDeployBlueGreenApplicationTarget{
//   				LogicalId: jsii.String("logicalId"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ServiceRole: jsii.String("serviceRole"),
//
//   	// the properties below are optional
//   	AdditionalOptions: &CfnCodeDeployBlueGreenAdditionalOptions{
//   		TerminationWaitTimeInMinutes: jsii.Number(123),
//   	},
//   	LifecycleEventHooks: &CfnCodeDeployBlueGreenLifecycleEventHooks{
//   		AfterAllowTestTraffic: jsii.String("afterAllowTestTraffic"),
//   		AfterAllowTraffic: jsii.String("afterAllowTraffic"),
//   		AfterInstall: jsii.String("afterInstall"),
//   		BeforeAllowTraffic: jsii.String("beforeAllowTraffic"),
//   		BeforeInstall: jsii.String("beforeInstall"),
//   	},
//   	TrafficRoutingConfig: &CfnTrafficRoutingConfig{
//   		Type: cdk.CfnTrafficRoutingType_ALL_AT_ONCE,
//
//   		// the properties below are optional
//   		TimeBasedCanary: &CfnTrafficRoutingTimeBasedCanary{
//   			BakeTimeMins: jsii.Number(123),
//   			StepPercentage: jsii.Number(123),
//   		},
//   		TimeBasedLinear: &CfnTrafficRoutingTimeBasedLinear{
//   			BakeTimeMins: jsii.Number(123),
//   			StepPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnCodeDeployBlueGreenHookProps struct {
	// Properties of the Amazon ECS applications being deployed.
	Applications *[]*CfnCodeDeployBlueGreenApplication `field:"required" json:"applications" yaml:"applications"`
	// The IAM Role for CloudFormation to use to perform blue-green deployments.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Additional options for the blue/green deployment.
	// Default: - no additional options.
	//
	AdditionalOptions *CfnCodeDeployBlueGreenAdditionalOptions `field:"optional" json:"additionalOptions" yaml:"additionalOptions"`
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda `CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic`
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	// Default: - no lifecycle event hooks.
	//
	LifecycleEventHooks *CfnCodeDeployBlueGreenLifecycleEventHooks `field:"optional" json:"lifecycleEventHooks" yaml:"lifecycleEventHooks"`
	// Traffic routing configuration settings.
	// Default: - time-based canary traffic shifting, with a 15% step percentage and a five minute bake time.
	//
	TrafficRoutingConfig *CfnTrafficRoutingConfig `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
}


// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Construction properties of {@link CfnCodeDeployBlueGreenHook}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenHookProps := &cfnCodeDeployBlueGreenHookProps{
//   	applications: []cfnCodeDeployBlueGreenApplication{
//   		&cfnCodeDeployBlueGreenApplication{
//   			ecsAttributes: &cfnCodeDeployBlueGreenEcsAttributes{
//   				taskDefinitions: []*string{
//   					jsii.String("taskDefinitions"),
//   				},
//   				taskSets: []*string{
//   					jsii.String("taskSets"),
//   				},
//   				trafficRouting: &cfnTrafficRouting{
//   					prodTrafficRoute: &cfnTrafficRoute{
//   						logicalId: jsii.String("logicalId"),
//   						type: jsii.String("type"),
//   					},
//   					targetGroups: []*string{
//   						jsii.String("targetGroups"),
//   					},
//   					testTrafficRoute: &cfnTrafficRoute{
//   						logicalId: jsii.String("logicalId"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			target: &cfnCodeDeployBlueGreenApplicationTarget{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//
//   	// the properties below are optional
//   	additionalOptions: &cfnCodeDeployBlueGreenAdditionalOptions{
//   		terminationWaitTimeInMinutes: jsii.Number(123),
//   	},
//   	lifecycleEventHooks: &cfnCodeDeployBlueGreenLifecycleEventHooks{
//   		afterAllowTestTraffic: jsii.String("afterAllowTestTraffic"),
//   		afterAllowTraffic: jsii.String("afterAllowTraffic"),
//   		afterInstall: jsii.String("afterInstall"),
//   		beforeAllowTraffic: jsii.String("beforeAllowTraffic"),
//   		beforeInstall: jsii.String("beforeInstall"),
//   	},
//   	trafficRoutingConfig: &cfnTrafficRoutingConfig{
//   		type: cdk.cfnTrafficRoutingType_ALL_AT_ONCE,
//
//   		// the properties below are optional
//   		timeBasedCanary: &cfnTrafficRoutingTimeBasedCanary{
//   			bakeTimeMins: jsii.Number(123),
//   			stepPercentage: jsii.Number(123),
//   		},
//   		timeBasedLinear: &cfnTrafficRoutingTimeBasedLinear{
//   			bakeTimeMins: jsii.Number(123),
//   			stepPercentage: jsii.Number(123),
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
	AdditionalOptions *CfnCodeDeployBlueGreenAdditionalOptions `field:"optional" json:"additionalOptions" yaml:"additionalOptions"`
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda {@link CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic}
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	LifecycleEventHooks *CfnCodeDeployBlueGreenLifecycleEventHooks `field:"optional" json:"lifecycleEventHooks" yaml:"lifecycleEventHooks"`
	// Traffic routing configuration settings.
	TrafficRoutingConfig *CfnTrafficRoutingConfig `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
}


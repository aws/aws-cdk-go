// An experiment to bundle the entire CDK into a single module
package awscdk


// Construction properties of {@link CfnCodeDeployBlueGreenHook}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
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
//   		type: monocdk.cfnTrafficRoutingType_ALL_AT_ONCE,
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
// Experimental.
type CfnCodeDeployBlueGreenHookProps struct {
	// Properties of the Amazon ECS applications being deployed.
	// Experimental.
	Applications *[]*CfnCodeDeployBlueGreenApplication `field:"required" json:"applications" yaml:"applications"`
	// The IAM Role for CloudFormation to use to perform blue-green deployments.
	// Experimental.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Additional options for the blue/green deployment.
	// Experimental.
	AdditionalOptions *CfnCodeDeployBlueGreenAdditionalOptions `field:"optional" json:"additionalOptions" yaml:"additionalOptions"`
	// Use lifecycle event hooks to specify a Lambda function that CodeDeploy can call to validate a deployment.
	//
	// You can use the same function or a different one for deployment lifecycle events.
	// Following completion of the validation tests,
	// the Lambda {@link CfnCodeDeployBlueGreenLifecycleEventHooks.afterAllowTraffic}
	// function calls back CodeDeploy and delivers a result of 'Succeeded' or 'Failed'.
	// Experimental.
	LifecycleEventHooks *CfnCodeDeployBlueGreenLifecycleEventHooks `field:"optional" json:"lifecycleEventHooks" yaml:"lifecycleEventHooks"`
	// Traffic routing configuration settings.
	// Experimental.
	TrafficRoutingConfig *CfnTrafficRoutingConfig `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
}


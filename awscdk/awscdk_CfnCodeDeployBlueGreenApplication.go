// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The application actually being deployed.
//
// Type of the {@link CfnCodeDeployBlueGreenHookProps.applications} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenApplication := &cfnCodeDeployBlueGreenApplication{
//   	ecsAttributes: &cfnCodeDeployBlueGreenEcsAttributes{
//   		taskDefinitions: []*string{
//   			jsii.String("taskDefinitions"),
//   		},
//   		taskSets: []*string{
//   			jsii.String("taskSets"),
//   		},
//   		trafficRouting: &cfnTrafficRouting{
//   			prodTrafficRoute: &cfnTrafficRoute{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   			targetGroups: []*string{
//   				jsii.String("targetGroups"),
//   			},
//   			testTrafficRoute: &cfnTrafficRoute{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	target: &cfnCodeDeployBlueGreenApplicationTarget{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnCodeDeployBlueGreenApplication struct {
	// The detailed attributes of the deployed target.
	EcsAttributes *CfnCodeDeployBlueGreenEcsAttributes `field:"required" json:"ecsAttributes" yaml:"ecsAttributes"`
	// The target that is being deployed.
	Target *CfnCodeDeployBlueGreenApplicationTarget `field:"required" json:"target" yaml:"target"`
}


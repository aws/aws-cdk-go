// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The attributes of the ECS Service being deployed.
//
// Type of the {@link CfnCodeDeployBlueGreenApplication.ecsAttributes} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenEcsAttributes := &cfnCodeDeployBlueGreenEcsAttributes{
//   	taskDefinitions: []*string{
//   		jsii.String("taskDefinitions"),
//   	},
//   	taskSets: []*string{
//   		jsii.String("taskSets"),
//   	},
//   	trafficRouting: &cfnTrafficRouting{
//   		prodTrafficRoute: &cfnTrafficRoute{
//   			logicalId: jsii.String("logicalId"),
//   			type: jsii.String("type"),
//   		},
//   		targetGroups: []*string{
//   			jsii.String("targetGroups"),
//   		},
//   		testTrafficRoute: &cfnTrafficRoute{
//   			logicalId: jsii.String("logicalId"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnCodeDeployBlueGreenEcsAttributes struct {
	// The logical IDs of the blue and green, respectively, AWS::ECS::TaskDefinition task definitions.
	TaskDefinitions *[]*string `field:"required" json:"taskDefinitions" yaml:"taskDefinitions"`
	// The logical IDs of the blue and green, respectively, AWS::ECS::TaskSet task sets.
	TaskSets *[]*string `field:"required" json:"taskSets" yaml:"taskSets"`
	// The traffic routing configuration.
	TrafficRouting *CfnTrafficRouting `field:"required" json:"trafficRouting" yaml:"trafficRouting"`
}


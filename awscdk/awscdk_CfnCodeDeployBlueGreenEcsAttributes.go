// An experiment to bundle the entire CDK into a single module
package awscdk


// The attributes of the ECS Service being deployed.
//
// Type of the {@link CfnCodeDeployBlueGreenApplication.ecsAttributes} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenEcsAttributes := &CfnCodeDeployBlueGreenEcsAttributes{
//   	TaskDefinitions: []*string{
//   		jsii.String("taskDefinitions"),
//   	},
//   	TaskSets: []*string{
//   		jsii.String("taskSets"),
//   	},
//   	TrafficRouting: &CfnTrafficRouting{
//   		ProdTrafficRoute: &CfnTrafficRoute{
//   			LogicalId: jsii.String("logicalId"),
//   			Type: jsii.String("type"),
//   		},
//   		TargetGroups: []*string{
//   			jsii.String("targetGroups"),
//   		},
//   		TestTrafficRoute: &CfnTrafficRoute{
//   			LogicalId: jsii.String("logicalId"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type CfnCodeDeployBlueGreenEcsAttributes struct {
	// The logical IDs of the blue and green, respectively, AWS::ECS::TaskDefinition task definitions.
	// Experimental.
	TaskDefinitions *[]*string `field:"required" json:"taskDefinitions" yaml:"taskDefinitions"`
	// The logical IDs of the blue and green, respectively, AWS::ECS::TaskSet task sets.
	// Experimental.
	TaskSets *[]*string `field:"required" json:"taskSets" yaml:"taskSets"`
	// The traffic routing configuration.
	// Experimental.
	TrafficRouting *CfnTrafficRouting `field:"required" json:"trafficRouting" yaml:"trafficRouting"`
}


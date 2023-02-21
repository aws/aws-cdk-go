package awsssm


// Properties for defining a `CfnMaintenanceWindowTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowTargetProps := &cfnMaintenanceWindowTargetProps{
//   	resourceType: jsii.String("resourceType"),
//   	targets: []interface{}{
//   		&targetsProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	ownerInformation: jsii.String("ownerInformation"),
//   }
//
type CfnMaintenanceWindowTargetProps struct {
	// The type of target that is being registered with the maintenance window.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The targets to register with the maintenance window.
	//
	// In other words, the instances to run commands on when the maintenance window runs.
	//
	// You must specify targets by using the `WindowTargetIds` parameter.
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The ID of the maintenance window to register the target with.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// A description for the target.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the maintenance window target.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window.
	OwnerInformation *string `field:"optional" json:"ownerInformation" yaml:"ownerInformation"`
}


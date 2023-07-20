package awsssm


// Properties for defining a `CfnMaintenanceWindowTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowTargetProps := &CfnMaintenanceWindowTargetProps{
//   	ResourceType: jsii.String("resourceType"),
//   	Targets: []interface{}{
//   		&TargetsProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	WindowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OwnerInformation: jsii.String("ownerInformation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html
//
type CfnMaintenanceWindowTargetProps struct {
	// The type of target that is being registered with the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html#cfn-ssm-maintenancewindowtarget-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The targets to register with the maintenance window.
	//
	// In other words, the instances to run commands on when the maintenance window runs.
	//
	// You must specify targets by using the `WindowTargetIds` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html#cfn-ssm-maintenancewindowtarget-targets
	//
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The ID of the maintenance window to register the target with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html#cfn-ssm-maintenancewindowtarget-windowid
	//
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// A description for the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html#cfn-ssm-maintenancewindowtarget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the maintenance window target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html#cfn-ssm-maintenancewindowtarget-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html#cfn-ssm-maintenancewindowtarget-ownerinformation
	//
	OwnerInformation *string `field:"optional" json:"ownerInformation" yaml:"ownerInformation"`
}


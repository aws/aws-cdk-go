package awsmediaconnect


// The maintenance setting of a flow.
//
// MediaConnect routinely performs maintenance on underlying systems for security, reliability, and operational performance. The maintenance activities include actions such as patching the operating system, updating drivers, or installing software and patches.
//
// You can select the day and time that maintenance events occur. This is called a maintenance window and is used every time a maintenance event is required. To change the day and time, you can edit the maintenance window using `MaintenanceDay` and `MaintenanceStartHour` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceProperty := &MaintenanceProperty{
//   	MaintenanceDay: jsii.String("maintenanceDay"),
//   	MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-maintenance.html
//
type CfnFlow_MaintenanceProperty struct {
	// A day of a week when the maintenance will happen.
	//
	// Use Monday/Tuesday/Wednesday/Thursday/Friday/Saturday/Sunday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-maintenance.html#cfn-mediaconnect-flow-maintenance-maintenanceday
	//
	MaintenanceDay *string `field:"required" json:"maintenanceDay" yaml:"maintenanceDay"`
	// UTC time when the maintenance will happen.
	//
	// Use 24-hour HH:MM format. Minutes must be 00. Example: 13:00. The default value is 02:00.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-maintenance.html#cfn-mediaconnect-flow-maintenance-maintenancestarthour
	//
	MaintenanceStartHour *string `field:"required" json:"maintenanceStartHour" yaml:"maintenanceStartHour"`
}


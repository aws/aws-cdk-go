package mixinsawsmediaconnect


// The maintenance setting of a flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   maintenanceProperty := &MaintenanceProperty{
//   	MaintenanceDay: jsii.String("maintenanceDay"),
//   	MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-maintenance.html
//
type CfnFlowPropsMixin_MaintenanceProperty struct {
	// A day of a week when the maintenance will happen.
	//
	// Use Monday/Tuesday/Wednesday/Thursday/Friday/Saturday/Sunday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-maintenance.html#cfn-mediaconnect-flow-maintenance-maintenanceday
	//
	MaintenanceDay *string `field:"optional" json:"maintenanceDay" yaml:"maintenanceDay"`
	// UTC time when the maintenance will happen.
	//
	// Use 24-hour HH:MM format. Minutes must be 00. Example: 13:00. The default value is 02:00.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-maintenance.html#cfn-mediaconnect-flow-maintenance-maintenancestarthour
	//
	MaintenanceStartHour *string `field:"optional" json:"maintenanceStartHour" yaml:"maintenanceStartHour"`
}


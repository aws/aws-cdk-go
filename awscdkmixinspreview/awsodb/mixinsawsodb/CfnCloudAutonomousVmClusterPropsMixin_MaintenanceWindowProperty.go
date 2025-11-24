package mixinsawsodb


// The scheduling details for the maintenance window.
//
// Patching and system updates take place during the maintenance window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   maintenanceWindowProperty := &MaintenanceWindowProperty{
//   	DaysOfWeek: []*string{
//   		jsii.String("daysOfWeek"),
//   	},
//   	HoursOfDay: []interface{}{
//   		jsii.Number(123),
//   	},
//   	LeadTimeInWeeks: jsii.Number(123),
//   	Months: []*string{
//   		jsii.String("months"),
//   	},
//   	Preference: jsii.String("preference"),
//   	WeeksOfMonth: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html
//
type CfnCloudAutonomousVmClusterPropsMixin_MaintenanceWindowProperty struct {
	// The days of the week when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow-daysofweek
	//
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The hours of the day when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow-hoursofday
	//
	HoursOfDay interface{} `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// The lead time in weeks before the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow-leadtimeinweeks
	//
	LeadTimeInWeeks *float64 `field:"optional" json:"leadTimeInWeeks" yaml:"leadTimeInWeeks"`
	// The months when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow-months
	//
	Months *[]*string `field:"optional" json:"months" yaml:"months"`
	// The preference for the maintenance window scheduling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow-preference
	//
	Preference *string `field:"optional" json:"preference" yaml:"preference"`
	// The weeks of the month when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow-weeksofmonth
	//
	WeeksOfMonth interface{} `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}


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
//   	CustomActionTimeoutInMins: jsii.Number(123),
//   	DaysOfWeek: []*string{
//   		jsii.String("daysOfWeek"),
//   	},
//   	HoursOfDay: []interface{}{
//   		jsii.Number(123),
//   	},
//   	IsCustomActionTimeoutEnabled: jsii.Boolean(false),
//   	LeadTimeInWeeks: jsii.Number(123),
//   	Months: []*string{
//   		jsii.String("months"),
//   	},
//   	PatchingMode: jsii.String("patchingMode"),
//   	Preference: jsii.String("preference"),
//   	WeeksOfMonth: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html
//
type CfnCloudExadataInfrastructurePropsMixin_MaintenanceWindowProperty struct {
	// The custom action timeout in minutes for the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-customactiontimeoutinmins
	//
	CustomActionTimeoutInMins *float64 `field:"optional" json:"customActionTimeoutInMins" yaml:"customActionTimeoutInMins"`
	// The days of the week when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-daysofweek
	//
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The hours of the day when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-hoursofday
	//
	HoursOfDay interface{} `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// Indicates whether custom action timeout is enabled for the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-iscustomactiontimeoutenabled
	//
	IsCustomActionTimeoutEnabled interface{} `field:"optional" json:"isCustomActionTimeoutEnabled" yaml:"isCustomActionTimeoutEnabled"`
	// The lead time in weeks before the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-leadtimeinweeks
	//
	LeadTimeInWeeks *float64 `field:"optional" json:"leadTimeInWeeks" yaml:"leadTimeInWeeks"`
	// The months when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-months
	//
	Months *[]*string `field:"optional" json:"months" yaml:"months"`
	// The patching mode for the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-patchingmode
	//
	PatchingMode *string `field:"optional" json:"patchingMode" yaml:"patchingMode"`
	// The preference for the maintenance window scheduling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-preference
	//
	Preference *string `field:"optional" json:"preference" yaml:"preference"`
	// The weeks of the month when maintenance can be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html#cfn-odb-cloudexadatainfrastructure-maintenancewindow-weeksofmonth
	//
	WeeksOfMonth interface{} `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}


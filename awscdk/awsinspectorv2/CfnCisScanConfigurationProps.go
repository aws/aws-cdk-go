package awsinspectorv2


// Properties for defining a `CfnCisScanConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var oneTime interface{}
//
//   cfnCisScanConfigurationProps := &CfnCisScanConfigurationProps{
//   	ScanName: jsii.String("scanName"),
//   	Schedule: &ScheduleProperty{
//   		Daily: &DailyScheduleProperty{
//   			StartTime: &TimeProperty{
//   				TimeOfDay: jsii.String("timeOfDay"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   		},
//   		Monthly: &MonthlyScheduleProperty{
//   			Day: jsii.String("day"),
//   			StartTime: &TimeProperty{
//   				TimeOfDay: jsii.String("timeOfDay"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   		},
//   		OneTime: oneTime,
//   		Weekly: &WeeklyScheduleProperty{
//   			Days: []*string{
//   				jsii.String("days"),
//   			},
//   			StartTime: &TimeProperty{
//   				TimeOfDay: jsii.String("timeOfDay"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   		},
//   	},
//   	SecurityLevel: jsii.String("securityLevel"),
//   	Targets: &CisTargetsProperty{
//   		AccountIds: []*string{
//   			jsii.String("accountIds"),
//   		},
//   		TargetResourceTags: map[string][]*string{
//   			"targetResourceTagsKey": []*string{
//   				jsii.String("targetResourceTags"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html
//
type CfnCisScanConfigurationProps struct {
	// The name of the CIS scan configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html#cfn-inspectorv2-cisscanconfiguration-scanname
	//
	ScanName *string `field:"required" json:"scanName" yaml:"scanName"`
	// The CIS scan configuration's schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html#cfn-inspectorv2-cisscanconfiguration-schedule
	//
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// The CIS scan configuration's CIS Benchmark level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html#cfn-inspectorv2-cisscanconfiguration-securitylevel
	//
	SecurityLevel *string `field:"required" json:"securityLevel" yaml:"securityLevel"`
	// The CIS scan configuration's targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html#cfn-inspectorv2-cisscanconfiguration-targets
	//
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The CIS scan configuration's tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html#cfn-inspectorv2-cisscanconfiguration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


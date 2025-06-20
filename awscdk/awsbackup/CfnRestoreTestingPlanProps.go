package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRestoreTestingPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRestoreTestingPlanProps := &CfnRestoreTestingPlanProps{
//   	RecoveryPointSelection: &RestoreTestingRecoveryPointSelectionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		IncludeVaults: []*string{
//   			jsii.String("includeVaults"),
//   		},
//   		RecoveryPointTypes: []*string{
//   			jsii.String("recoveryPointTypes"),
//   		},
//
//   		// the properties below are optional
//   		ExcludeVaults: []*string{
//   			jsii.String("excludeVaults"),
//   		},
//   		SelectionWindowDays: jsii.Number(123),
//   	},
//   	RestoreTestingPlanName: jsii.String("restoreTestingPlanName"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	StartWindowHours: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html
//
type CfnRestoreTestingPlanProps struct {
	// The specified criteria to assign a set of resources, such as recovery point types or backup vaults.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html#cfn-backup-restoretestingplan-recoverypointselection
	//
	RecoveryPointSelection interface{} `field:"required" json:"recoveryPointSelection" yaml:"recoveryPointSelection"`
	// The RestoreTestingPlanName is a unique string that is the name of the restore testing plan.
	//
	// This cannot be changed after creation, and it must consist of only alphanumeric characters and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html#cfn-backup-restoretestingplan-restoretestingplanname
	//
	RestoreTestingPlanName *string `field:"required" json:"restoreTestingPlanName" yaml:"restoreTestingPlanName"`
	// A CRON expression in specified timezone when a restore testing plan is executed.
	//
	// When no CRON expression is provided, AWS Backup will use the default expression `cron(0 5 ? * * *)` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html#cfn-backup-restoretestingplan-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Optional.
	//
	// This is the timezone in which the schedule expression is set. By default, ScheduleExpressions are in UTC. You can modify this to a specified timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html#cfn-backup-restoretestingplan-scheduleexpressiontimezone
	//
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// Defaults to 24 hours.
	//
	// A value in hours after a restore test is scheduled before a job will be canceled if it doesn't start successfully. This value is optional. If this value is included, this parameter has a maximum value of 168 hours (one week).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html#cfn-backup-restoretestingplan-startwindowhours
	//
	StartWindowHours *float64 `field:"optional" json:"startWindowHours" yaml:"startWindowHours"`
	// Optional tags to include.
	//
	// A tag is a key-value pair you can use to manage, filter, and search for your resources. Allowed characters include UTF-8 letters,numbers, spaces, and the following characters: `+ - = . _ : /.`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-restoretestingplan.html#cfn-backup-restoretestingplan-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


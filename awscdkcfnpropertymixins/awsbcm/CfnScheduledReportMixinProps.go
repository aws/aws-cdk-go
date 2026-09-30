package awsbcm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnScheduledReportPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnScheduledReportMixinProps := &CfnScheduledReportMixinProps{
//   	DashboardArn: jsii.String("dashboardArn"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		ScheduleExpressionTimeZone: jsii.String("scheduleExpressionTimeZone"),
//   		SchedulePeriod: &SchedulePeriodProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	ScheduledReportExecutionRoleArn: jsii.String("scheduledReportExecutionRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WidgetDateRangeOverride: &DateTimeRangeProperty{
//   		EndTime: &DateTimeValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   		StartTime: &DateTimeValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WidgetIds: []*string{
//   		jsii.String("widgetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html
//
type CfnScheduledReportMixinProps struct {
	// The ARN of the dashboard associated with the scheduled report.
	//
	// Managed dashboards cannot be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-dashboardarn
	//
	DashboardArn *string `field:"optional" json:"dashboardArn" yaml:"dashboardArn"`
	// A description of the scheduled report's purpose or contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the scheduled report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Defines when and how often a scheduled report runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-scheduleconfig
	//
	ScheduleConfig interface{} `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
	// The ARN of the IAM role that the scheduled report uses to execute.
	//
	// AWS Billing and Cost Management Dashboards assumes this IAM role while executing the scheduled report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-scheduledreportexecutionrolearn
	//
	ScheduledReportExecutionRoleArn *string `field:"optional" json:"scheduledReportExecutionRoleArn" yaml:"scheduledReportExecutionRoleArn"`
	// The tags applied to the scheduled report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An absolute or relative date range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-widgetdaterangeoverride
	//
	WidgetDateRangeOverride interface{} `field:"optional" json:"widgetDateRangeOverride" yaml:"widgetDateRangeOverride"`
	// The list of widget identifiers included in the scheduled report.
	//
	// If not specified, all widgets in the dashboard are included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html#cfn-bcm-scheduledreport-widgetids
	//
	WidgetIds *[]*string `field:"optional" json:"widgetIds" yaml:"widgetIds"`
}


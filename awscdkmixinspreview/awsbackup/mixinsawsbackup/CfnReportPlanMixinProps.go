package mixinsawsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnReportPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var reportDeliveryChannel interface{}
//   var reportSetting interface{}
//
//   cfnReportPlanMixinProps := &CfnReportPlanMixinProps{
//   	ReportDeliveryChannel: reportDeliveryChannel,
//   	ReportPlanDescription: jsii.String("reportPlanDescription"),
//   	ReportPlanName: jsii.String("reportPlanName"),
//   	ReportPlanTags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ReportSetting: reportSetting,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html
//
type CfnReportPlanMixinProps struct {
	// Contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html#cfn-backup-reportplan-reportdeliverychannel
	//
	ReportDeliveryChannel interface{} `field:"optional" json:"reportDeliveryChannel" yaml:"reportDeliveryChannel"`
	// An optional description of the report plan with a maximum 1,024 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html#cfn-backup-reportplan-reportplandescription
	//
	ReportPlanDescription *string `field:"optional" json:"reportPlanDescription" yaml:"reportPlanDescription"`
	// The unique name of the report plan.
	//
	// This name is between 1 and 256 characters starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html#cfn-backup-reportplan-reportplanname
	//
	ReportPlanName *string `field:"optional" json:"reportPlanName" yaml:"reportPlanName"`
	// The tags to assign to your report plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html#cfn-backup-reportplan-reportplantags
	//
	ReportPlanTags *[]*awscdk.CfnTag `field:"optional" json:"reportPlanTags" yaml:"reportPlanTags"`
	// Identifies the report template for the report. Reports are built using a report template. The report templates are:.
	//
	// `RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`
	//
	// If the report template is `RESOURCE_COMPLIANCE_REPORT` or `CONTROL_COMPLIANCE_REPORT` , this API resource also describes the report coverage by AWS Regions and frameworks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html#cfn-backup-reportplan-reportsetting
	//
	ReportSetting interface{} `field:"optional" json:"reportSetting" yaml:"reportSetting"`
}


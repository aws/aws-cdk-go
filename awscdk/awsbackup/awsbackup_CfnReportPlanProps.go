package awsbackup


// Properties for defining a `CfnReportPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var reportDeliveryChannel interface{}
//   var reportSetting interface{}
//
//   cfnReportPlanProps := &cfnReportPlanProps{
//   	reportDeliveryChannel: reportDeliveryChannel,
//   	reportSetting: reportSetting,
//
//   	// the properties below are optional
//   	reportPlanDescription: jsii.String("reportPlanDescription"),
//   	reportPlanName: jsii.String("reportPlanName"),
//   	reportPlanTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnReportPlanProps struct {
	// Contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
	ReportDeliveryChannel interface{} `field:"required" json:"reportDeliveryChannel" yaml:"reportDeliveryChannel"`
	// Identifies the report template for the report. Reports are built using a report template. The report templates are:.
	//
	// `RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`
	//
	// If the report template is `RESOURCE_COMPLIANCE_REPORT` or `CONTROL_COMPLIANCE_REPORT` , this API resource also describes the report coverage by AWS Regions and frameworks.
	ReportSetting interface{} `field:"required" json:"reportSetting" yaml:"reportSetting"`
	// An optional description of the report plan with a maximum 1,024 characters.
	ReportPlanDescription *string `field:"optional" json:"reportPlanDescription" yaml:"reportPlanDescription"`
	// The unique name of the report plan.
	//
	// This name is between 1 and 256 characters starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	ReportPlanName *string `field:"optional" json:"reportPlanName" yaml:"reportPlanName"`
	// A list of tags to tag your report plan.
	ReportPlanTags interface{} `field:"optional" json:"reportPlanTags" yaml:"reportPlanTags"`
}


package awsbackup


// Contains detailed information about a report setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportSettingProperty := &ReportSettingProperty{
//   	ReportTemplate: jsii.String("reportTemplate"),
//
//   	// the properties below are optional
//   	Accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	FrameworkArns: []*string{
//   		jsii.String("frameworkArns"),
//   	},
//   	OrganizationUnits: []*string{
//   		jsii.String("organizationUnits"),
//   	},
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportsetting.html
//
type CfnReportPlan_ReportSettingProperty struct {
	// Identifies the report template for the report. Reports are built using a report template. The report templates are:.
	//
	// `RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportsetting.html#cfn-backup-reportplan-reportsetting-reporttemplate
	//
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// These are the accounts to be included in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportsetting.html#cfn-backup-reportplan-reportsetting-accounts
	//
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The Amazon Resource Names (ARNs) of the frameworks a report covers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportsetting.html#cfn-backup-reportplan-reportsetting-frameworkarns
	//
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// These are the Organizational Units to be included in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportsetting.html#cfn-backup-reportplan-reportsetting-organizationunits
	//
	OrganizationUnits *[]*string `field:"optional" json:"organizationUnits" yaml:"organizationUnits"`
	// These are the Regions to be included in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportsetting.html#cfn-backup-reportplan-reportsetting-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


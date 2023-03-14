package awsbackup


// Contains detailed information about a report setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportSettingProperty := &reportSettingProperty{
//   	reportTemplate: jsii.String("reportTemplate"),
//
//   	// the properties below are optional
//   	accounts: []*string{
//   		jsii.String("accounts"),
//   	},
//   	frameworkArns: []*string{
//   		jsii.String("frameworkArns"),
//   	},
//   	organizationUnits: []*string{
//   		jsii.String("organizationUnits"),
//   	},
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
type CfnReportPlan_ReportSettingProperty struct {
	// Identifies the report template for the report. Reports are built using a report template. The report templates are:.
	//
	// `RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`.
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// These are the accounts to be included in the report.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The Amazon Resource Names (ARNs) of the frameworks a report covers.
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// These are the Organizational Units to be included in the report.
	OrganizationUnits *[]*string `field:"optional" json:"organizationUnits" yaml:"organizationUnits"`
	// These are the Regions to be included in the report.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


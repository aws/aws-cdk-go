package awsbackup


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
	// `CfnReportPlan.ReportSettingProperty.ReportTemplate`.
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// `CfnReportPlan.ReportSettingProperty.Accounts`.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// `CfnReportPlan.ReportSettingProperty.FrameworkArns`.
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// `CfnReportPlan.ReportSettingProperty.OrganizationUnits`.
	OrganizationUnits *[]*string `field:"optional" json:"organizationUnits" yaml:"organizationUnits"`
	// `CfnReportPlan.ReportSettingProperty.Regions`.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


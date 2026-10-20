package awslicensemanager


// Details of the license configurations and asset groups that this generator reports on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportContextProperty := &ReportContextProperty{
//   	LicenseAssetGroupArns: []*string{
//   		jsii.String("licenseAssetGroupArns"),
//   	},
//   	LicenseConfigurationArns: []*string{
//   		jsii.String("licenseConfigurationArns"),
//   	},
//   	ReportEndDate: jsii.String("reportEndDate"),
//   	ReportStartDate: jsii.String("reportStartDate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportcontext.html
//
type CfnReportGenerator_ReportContextProperty struct {
	// Amazon Resource Names (ARNs) of the license asset groups to include in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportcontext.html#cfn-licensemanager-reportgenerator-reportcontext-licenseassetgrouparns
	//
	LicenseAssetGroupArns *[]*string `field:"optional" json:"licenseAssetGroupArns" yaml:"licenseAssetGroupArns"`
	// Amazon Resource Names (ARNs) of the license configurations that this generator reports on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportcontext.html#cfn-licensemanager-reportgenerator-reportcontext-licenseconfigurationarns
	//
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// End date for the report data collection period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportcontext.html#cfn-licensemanager-reportgenerator-reportcontext-reportenddate
	//
	ReportEndDate *string `field:"optional" json:"reportEndDate" yaml:"reportEndDate"`
	// Start date for the report data collection period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-reportgenerator-reportcontext.html#cfn-licensemanager-reportgenerator-reportcontext-reportstartdate
	//
	ReportStartDate *string `field:"optional" json:"reportStartDate" yaml:"reportStartDate"`
}


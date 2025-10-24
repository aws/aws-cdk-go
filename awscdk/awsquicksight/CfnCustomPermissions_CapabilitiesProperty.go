package awsquicksight


// A set of actions that correspond to Amazon Quick Sight permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilitiesProperty := &CapabilitiesProperty{
//   	AddOrRunAnomalyDetectionForAnalyses: jsii.String("addOrRunAnomalyDetectionForAnalyses"),
//   	Analysis: jsii.String("analysis"),
//   	CreateAndUpdateDashboardEmailReports: jsii.String("createAndUpdateDashboardEmailReports"),
//   	CreateAndUpdateDatasets: jsii.String("createAndUpdateDatasets"),
//   	CreateAndUpdateDataSources: jsii.String("createAndUpdateDataSources"),
//   	CreateAndUpdateThemes: jsii.String("createAndUpdateThemes"),
//   	CreateAndUpdateThresholdAlerts: jsii.String("createAndUpdateThresholdAlerts"),
//   	CreateSharedFolders: jsii.String("createSharedFolders"),
//   	CreateSpiceDataset: jsii.String("createSpiceDataset"),
//   	Dashboard: jsii.String("dashboard"),
//   	ExportToCsv: jsii.String("exportToCsv"),
//   	ExportToCsvInScheduledReports: jsii.String("exportToCsvInScheduledReports"),
//   	ExportToExcel: jsii.String("exportToExcel"),
//   	ExportToExcelInScheduledReports: jsii.String("exportToExcelInScheduledReports"),
//   	ExportToPdf: jsii.String("exportToPdf"),
//   	ExportToPdfInScheduledReports: jsii.String("exportToPdfInScheduledReports"),
//   	IncludeContentInScheduledReportsEmail: jsii.String("includeContentInScheduledReportsEmail"),
//   	PrintReports: jsii.String("printReports"),
//   	RenameSharedFolders: jsii.String("renameSharedFolders"),
//   	ShareAnalyses: jsii.String("shareAnalyses"),
//   	ShareDashboards: jsii.String("shareDashboards"),
//   	ShareDatasets: jsii.String("shareDatasets"),
//   	ShareDataSources: jsii.String("shareDataSources"),
//   	SubscribeDashboardEmailReports: jsii.String("subscribeDashboardEmailReports"),
//   	ViewAccountSpiceCapacity: jsii.String("viewAccountSpiceCapacity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html
//
type CfnCustomPermissions_CapabilitiesProperty struct {
	// The ability to add or run anomaly detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-addorrunanomalydetectionforanalyses
	//
	AddOrRunAnomalyDetectionForAnalyses *string `field:"optional" json:"addOrRunAnomalyDetectionForAnalyses" yaml:"addOrRunAnomalyDetectionForAnalyses"`
	// The ability to perform analysis-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-analysis
	//
	Analysis *string `field:"optional" json:"analysis" yaml:"analysis"`
	// The ability to create and update email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatedashboardemailreports
	//
	CreateAndUpdateDashboardEmailReports *string `field:"optional" json:"createAndUpdateDashboardEmailReports" yaml:"createAndUpdateDashboardEmailReports"`
	// The ability to create and update datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatedatasets
	//
	CreateAndUpdateDatasets *string `field:"optional" json:"createAndUpdateDatasets" yaml:"createAndUpdateDatasets"`
	// The ability to create and update data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatedatasources
	//
	CreateAndUpdateDataSources *string `field:"optional" json:"createAndUpdateDataSources" yaml:"createAndUpdateDataSources"`
	// The ability to export to Create and Update themes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatethemes
	//
	CreateAndUpdateThemes *string `field:"optional" json:"createAndUpdateThemes" yaml:"createAndUpdateThemes"`
	// The ability to create and update threshold alerts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createandupdatethresholdalerts
	//
	CreateAndUpdateThresholdAlerts *string `field:"optional" json:"createAndUpdateThresholdAlerts" yaml:"createAndUpdateThresholdAlerts"`
	// The ability to create shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createsharedfolders
	//
	CreateSharedFolders *string `field:"optional" json:"createSharedFolders" yaml:"createSharedFolders"`
	// The ability to create a SPICE dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-createspicedataset
	//
	CreateSpiceDataset *string `field:"optional" json:"createSpiceDataset" yaml:"createSpiceDataset"`
	// The ability to perform dashboard-related actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-dashboard
	//
	Dashboard *string `field:"optional" json:"dashboard" yaml:"dashboard"`
	// The ability to export to CSV files from the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttocsv
	//
	ExportToCsv *string `field:"optional" json:"exportToCsv" yaml:"exportToCsv"`
	// The ability to export to CSV files in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttocsvinscheduledreports
	//
	ExportToCsvInScheduledReports *string `field:"optional" json:"exportToCsvInScheduledReports" yaml:"exportToCsvInScheduledReports"`
	// The ability to export to Excel files from the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttoexcel
	//
	ExportToExcel *string `field:"optional" json:"exportToExcel" yaml:"exportToExcel"`
	// The ability to export to Excel files in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttoexcelinscheduledreports
	//
	ExportToExcelInScheduledReports *string `field:"optional" json:"exportToExcelInScheduledReports" yaml:"exportToExcelInScheduledReports"`
	// The ability to export to PDF files from the UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttopdf
	//
	ExportToPdf *string `field:"optional" json:"exportToPdf" yaml:"exportToPdf"`
	// The ability to export to PDF files in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-exporttopdfinscheduledreports
	//
	ExportToPdfInScheduledReports *string `field:"optional" json:"exportToPdfInScheduledReports" yaml:"exportToPdfInScheduledReports"`
	// The ability to include content in scheduled email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-includecontentinscheduledreportsemail
	//
	IncludeContentInScheduledReportsEmail *string `field:"optional" json:"includeContentInScheduledReportsEmail" yaml:"includeContentInScheduledReportsEmail"`
	// The ability to print reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-printreports
	//
	PrintReports *string `field:"optional" json:"printReports" yaml:"printReports"`
	// The ability to rename shared folders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-renamesharedfolders
	//
	RenameSharedFolders *string `field:"optional" json:"renameSharedFolders" yaml:"renameSharedFolders"`
	// The ability to share analyses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-shareanalyses
	//
	ShareAnalyses *string `field:"optional" json:"shareAnalyses" yaml:"shareAnalyses"`
	// The ability to share dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedashboards
	//
	ShareDashboards *string `field:"optional" json:"shareDashboards" yaml:"shareDashboards"`
	// The ability to share datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatasets
	//
	ShareDatasets *string `field:"optional" json:"shareDatasets" yaml:"shareDatasets"`
	// The ability to share data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-sharedatasources
	//
	ShareDataSources *string `field:"optional" json:"shareDataSources" yaml:"shareDataSources"`
	// The ability to subscribe to email reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-subscribedashboardemailreports
	//
	SubscribeDashboardEmailReports *string `field:"optional" json:"subscribeDashboardEmailReports" yaml:"subscribeDashboardEmailReports"`
	// The ability to view account SPICE capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-custompermissions-capabilities.html#cfn-quicksight-custompermissions-capabilities-viewaccountspicecapacity
	//
	ViewAccountSpiceCapacity *string `field:"optional" json:"viewAccountSpiceCapacity" yaml:"viewAccountSpiceCapacity"`
}


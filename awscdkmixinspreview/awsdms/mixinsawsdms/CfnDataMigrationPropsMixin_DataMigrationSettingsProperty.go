package mixinsawsdms


// Options for configuring a data migration, including whether to enable CloudWatch logs, and the selection rules to use to include or exclude database objects from the migration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataMigrationSettingsProperty := &DataMigrationSettingsProperty{
//   	CloudwatchLogsEnabled: jsii.Boolean(false),
//   	NumberOfJobs: jsii.Number(123),
//   	SelectionRules: jsii.String("selectionRules"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-datamigrationsettings.html
//
type CfnDataMigrationPropsMixin_DataMigrationSettingsProperty struct {
	// Whether to enable CloudWatch logging for the data migration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-datamigrationsettings.html#cfn-dms-datamigration-datamigrationsettings-cloudwatchlogsenabled
	//
	CloudwatchLogsEnabled interface{} `field:"optional" json:"cloudwatchLogsEnabled" yaml:"cloudwatchLogsEnabled"`
	// The number of parallel jobs that trigger parallel threads to unload the tables from the source, and then load them to the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-datamigrationsettings.html#cfn-dms-datamigration-datamigrationsettings-numberofjobs
	//
	NumberOfJobs *float64 `field:"optional" json:"numberOfJobs" yaml:"numberOfJobs"`
	// A JSON-formatted string that defines what objects to include and exclude from the migration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-datamigrationsettings.html#cfn-dms-datamigration-datamigrationsettings-selectionrules
	//
	SelectionRules *string `field:"optional" json:"selectionRules" yaml:"selectionRules"`
}


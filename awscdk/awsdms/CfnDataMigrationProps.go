package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataMigration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataMigrationProps := &CfnDataMigrationProps{
//   	DataMigrationType: jsii.String("dataMigrationType"),
//   	MigrationProjectIdentifier: jsii.String("migrationProjectIdentifier"),
//   	ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//
//   	// the properties below are optional
//   	DataMigrationIdentifier: jsii.String("dataMigrationIdentifier"),
//   	DataMigrationName: jsii.String("dataMigrationName"),
//   	DataMigrationSettings: &DataMigrationSettingsProperty{
//   		CloudwatchLogsEnabled: jsii.Boolean(false),
//   		NumberOfJobs: jsii.Number(123),
//   		SelectionRules: jsii.String("selectionRules"),
//   	},
//   	SourceDataSettings: []interface{}{
//   		&SourceDataSettingsProperty{
//   			CdcStartPosition: jsii.String("cdcStartPosition"),
//   			CdcStartTime: jsii.String("cdcStartTime"),
//   			CdcStopTime: jsii.String("cdcStopTime"),
//   			SlotName: jsii.String("slotName"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html
//
type CfnDataMigrationProps struct {
	// Specifies whether the data migration is full-load only, change data capture (CDC) only, or full-load and CDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationtype
	//
	DataMigrationType *string `field:"required" json:"dataMigrationType" yaml:"dataMigrationType"`
	// The property describes an identifier for the migration project.
	//
	// It is used for describing/deleting/modifying can be name/arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-migrationprojectidentifier
	//
	MigrationProjectIdentifier *string `field:"required" json:"migrationProjectIdentifier" yaml:"migrationProjectIdentifier"`
	// The IAM role that the data migration uses to access AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-serviceaccessrolearn
	//
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The property describes an ARN of the data migration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationidentifier
	//
	DataMigrationIdentifier *string `field:"optional" json:"dataMigrationIdentifier" yaml:"dataMigrationIdentifier"`
	// The user-friendly name for the data migration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationname
	//
	DataMigrationName *string `field:"optional" json:"dataMigrationName" yaml:"dataMigrationName"`
	// Specifies CloudWatch settings and selection rules for the data migration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-datamigrationsettings
	//
	DataMigrationSettings interface{} `field:"optional" json:"dataMigrationSettings" yaml:"dataMigrationSettings"`
	// Specifies information about the data migration's source data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-sourcedatasettings
	//
	SourceDataSettings interface{} `field:"optional" json:"sourceDataSettings" yaml:"sourceDataSettings"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-datamigration.html#cfn-dms-datamigration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


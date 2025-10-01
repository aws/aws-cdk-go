package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDatabase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatabaseProps := &CfnDatabaseProps{
//   	MasterDatabaseName: jsii.String("masterDatabaseName"),
//   	MasterUsername: jsii.String("masterUsername"),
//   	RelationalDatabaseBlueprintId: jsii.String("relationalDatabaseBlueprintId"),
//   	RelationalDatabaseBundleId: jsii.String("relationalDatabaseBundleId"),
//   	RelationalDatabaseName: jsii.String("relationalDatabaseName"),
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BackupRetention: jsii.Boolean(false),
//   	CaCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	RelationalDatabaseParameters: []interface{}{
//   		&RelationalDatabaseParameterProperty{
//   			AllowedValues: jsii.String("allowedValues"),
//   			ApplyMethod: jsii.String("applyMethod"),
//   			ApplyType: jsii.String("applyType"),
//   			DataType: jsii.String("dataType"),
//   			Description: jsii.String("description"),
//   			IsModifiable: jsii.Boolean(false),
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	RotateMasterUserPassword: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html
//
type CfnDatabaseProps struct {
	// The meaning of this parameter differs according to the database engine you use.
	//
	// *MySQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, no database is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-64 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in MySQL, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , and [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, a database named `postgres` is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-63 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in PostgreSQL, see the SQL Key Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-masterdatabasename
	//
	MasterDatabaseName *string `field:"required" json:"masterDatabaseName" yaml:"masterDatabaseName"`
	// The name for the primary user.
	//
	// *MySQL*
	//
	// Constraints:
	//
	// - Required for MySQL.
	// - Must be 1-16 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , or [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// Constraints:
	//
	// - Required for PostgreSQL.
	// - Must be 1-63 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-masterusername
	//
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// The blueprint ID for the database (for example, `mysql_8_0` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabaseblueprintid
	//
	RelationalDatabaseBlueprintId *string `field:"required" json:"relationalDatabaseBlueprintId" yaml:"relationalDatabaseBlueprintId"`
	// The bundle ID for the database (for example, `medium_1_0` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabasebundleid
	//
	RelationalDatabaseBundleId *string `field:"required" json:"relationalDatabaseBundleId" yaml:"relationalDatabaseBundleId"`
	// The name of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabasename
	//
	RelationalDatabaseName *string `field:"required" json:"relationalDatabaseName" yaml:"relationalDatabaseName"`
	// The Availability Zone for the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// A Boolean value indicating whether automated backup retention is enabled for the database.
	//
	// Data Import Mode is enabled when `BackupRetention` is set to `false` , and is disabled when `BackupRetention` is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-backupretention
	//
	BackupRetention interface{} `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// The certificate associated with the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-cacertificateidentifier
	//
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// The password for the primary user of the database.
	//
	// The password can include any printable ASCII character except the following: /, ", or.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-masteruserpassword
	//
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// The daily time range during which automated backups are created for the database (for example, `16:00-16:30` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-preferredbackupwindow
	//
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur for the database, formatted as follows: `ddd:hh24:mi-ddd:hh24:mi` .
	//
	// For example, `Tue:17:00-Tue:17:30` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A Boolean value indicating whether the database is accessible to anyone on the internet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An array of parameters for the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-relationaldatabaseparameters
	//
	RelationalDatabaseParameters interface{} `field:"optional" json:"relationalDatabaseParameters" yaml:"relationalDatabaseParameters"`
	// A Boolean value indicating whether to change the primary user password to a new, strong password generated by Lightsail .
	//
	// > The `RotateMasterUserPassword` and `MasterUserPassword` parameters cannot be used together in the same template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-rotatemasteruserpassword
	//
	RotateMasterUserPassword interface{} `field:"optional" json:"rotateMasterUserPassword" yaml:"rotateMasterUserPassword"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html#cfn-lightsail-database-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


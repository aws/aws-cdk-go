package awskendra


// Provides the configuration information to connect to a index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseConfigurationProperty := &DatabaseConfigurationProperty{
//   	ColumnConfiguration: &ColumnConfigurationProperty{
//   		ChangeDetectingColumns: []*string{
//   			jsii.String("changeDetectingColumns"),
//   		},
//   		DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   		DocumentIdColumnName: jsii.String("documentIdColumnName"),
//
//   		// the properties below are optional
//   		DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	ConnectionConfiguration: &ConnectionConfigurationProperty{
//   		DatabaseHost: jsii.String("databaseHost"),
//   		DatabaseName: jsii.String("databaseName"),
//   		DatabasePort: jsii.Number(123),
//   		SecretArn: jsii.String("secretArn"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	DatabaseEngineType: jsii.String("databaseEngineType"),
//
//   	// the properties below are optional
//   	AclConfiguration: &AclConfigurationProperty{
//   		AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   	},
//   	SqlConfiguration: &SqlConfigurationProperty{
//   		QueryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   	},
//   	VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataSource_DatabaseConfigurationProperty struct {
	// Information about where the index should get the document information from the database.
	ColumnConfiguration interface{} `field:"required" json:"columnConfiguration" yaml:"columnConfiguration"`
	// Configuration information that's required to connect to a database.
	ConnectionConfiguration interface{} `field:"required" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// The type of database engine that runs the database.
	DatabaseEngineType *string `field:"required" json:"databaseEngineType" yaml:"databaseEngineType"`
	// Information about the database column that provides information for user context filtering.
	AclConfiguration interface{} `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// Provides information about how Amazon Kendra uses quote marks around SQL identifiers when querying a database data source.
	SqlConfiguration interface{} `field:"optional" json:"sqlConfiguration" yaml:"sqlConfiguration"`
	// Provides information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}


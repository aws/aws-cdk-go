package awskendra


// Provides the configuration information to connect to a index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseConfigurationProperty := &databaseConfigurationProperty{
//   	columnConfiguration: &columnConfigurationProperty{
//   		changeDetectingColumns: []*string{
//   			jsii.String("changeDetectingColumns"),
//   		},
//   		documentDataColumnName: jsii.String("documentDataColumnName"),
//   		documentIdColumnName: jsii.String("documentIdColumnName"),
//
//   		// the properties below are optional
//   		documentTitleColumnName: jsii.String("documentTitleColumnName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	connectionConfiguration: &connectionConfigurationProperty{
//   		databaseHost: jsii.String("databaseHost"),
//   		databaseName: jsii.String("databaseName"),
//   		databasePort: jsii.Number(123),
//   		secretArn: jsii.String("secretArn"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	databaseEngineType: jsii.String("databaseEngineType"),
//
//   	// the properties below are optional
//   	aclConfiguration: &aclConfigurationProperty{
//   		allowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   	},
//   	sqlConfiguration: &sqlConfigurationProperty{
//   		queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   	},
//   	vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
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


package mixinsawskendra


// Provides the configuration information to an [Amazon Kendra supported database](https://docs.aws.amazon.com/kendra/latest/dg/data-source-database.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   databaseConfigurationProperty := &DatabaseConfigurationProperty{
//   	AclConfiguration: &AclConfigurationProperty{
//   		AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   	},
//   	ColumnConfiguration: &ColumnConfigurationProperty{
//   		ChangeDetectingColumns: []*string{
//   			jsii.String("changeDetectingColumns"),
//   		},
//   		DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   		DocumentIdColumnName: jsii.String("documentIdColumnName"),
//   		DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html
//
type CfnDataSourcePropsMixin_DatabaseConfigurationProperty struct {
	// Information about the database column that provides information for user context filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html#cfn-kendra-datasource-databaseconfiguration-aclconfiguration
	//
	AclConfiguration interface{} `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// Information about where the index should get the document information from the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html#cfn-kendra-datasource-databaseconfiguration-columnconfiguration
	//
	ColumnConfiguration interface{} `field:"optional" json:"columnConfiguration" yaml:"columnConfiguration"`
	// Configuration information that's required to connect to a database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html#cfn-kendra-datasource-databaseconfiguration-connectionconfiguration
	//
	ConnectionConfiguration interface{} `field:"optional" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// The type of database engine that runs the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html#cfn-kendra-datasource-databaseconfiguration-databaseenginetype
	//
	DatabaseEngineType *string `field:"optional" json:"databaseEngineType" yaml:"databaseEngineType"`
	// Provides information about how Amazon Kendra uses quote marks around SQL identifiers when querying a database data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html#cfn-kendra-datasource-databaseconfiguration-sqlconfiguration
	//
	SqlConfiguration interface{} `field:"optional" json:"sqlConfiguration" yaml:"sqlConfiguration"`
	// Provides information for connecting to an Amazon VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-databaseconfiguration.html#cfn-kendra-datasource-databaseconfiguration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}


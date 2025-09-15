package awscleanroomsml


// Defines the Glue data source and schema mapping information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetInputConfigProperty := &DatasetInputConfigProperty{
//   	DataSource: &DataSourceProperty{
//   		GlueDataSource: &GlueDataSourceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			CatalogId: jsii.String("catalogId"),
//   		},
//   	},
//   	Schema: []interface{}{
//   		&ColumnSchemaProperty{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnTypes: []*string{
//   				jsii.String("columnTypes"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasetinputconfig.html
//
type CfnTrainingDataset_DatasetInputConfigProperty struct {
	// A DataSource object that specifies the Glue data source for the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasetinputconfig.html#cfn-cleanroomsml-trainingdataset-datasetinputconfig-datasource
	//
	DataSource interface{} `field:"required" json:"dataSource" yaml:"dataSource"`
	// The schema information for the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasetinputconfig.html#cfn-cleanroomsml-trainingdataset-datasetinputconfig-schema
	//
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
}


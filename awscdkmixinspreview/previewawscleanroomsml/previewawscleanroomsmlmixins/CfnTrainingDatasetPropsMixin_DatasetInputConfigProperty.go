package previewawscleanroomsmlmixins


// Defines the Glue data source and schema mapping information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datasetInputConfigProperty := &DatasetInputConfigProperty{
//   	DataSource: &DataSourceProperty{
//   		GlueDataSource: &GlueDataSourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
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
type CfnTrainingDatasetPropsMixin_DatasetInputConfigProperty struct {
	// A DataSource object that specifies the Glue data source for the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasetinputconfig.html#cfn-cleanroomsml-trainingdataset-datasetinputconfig-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The schema information for the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasetinputconfig.html#cfn-cleanroomsml-trainingdataset-datasetinputconfig-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}


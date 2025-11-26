package previewawscleanroomsmlmixins


// Defines where the training dataset is located, what type of data it contains, and how to access the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datasetProperty := &DatasetProperty{
//   	InputConfig: &DatasetInputConfigProperty{
//   		DataSource: &DataSourceProperty{
//   			GlueDataSource: &GlueDataSourceProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				DatabaseName: jsii.String("databaseName"),
//   				TableName: jsii.String("tableName"),
//   			},
//   		},
//   		Schema: []interface{}{
//   			&ColumnSchemaProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ColumnTypes: []*string{
//   					jsii.String("columnTypes"),
//   				},
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-dataset.html
//
type CfnTrainingDatasetPropsMixin_DatasetProperty struct {
	// A DatasetInputConfig object that defines the data source and schema mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-dataset.html#cfn-cleanroomsml-trainingdataset-dataset-inputconfig
	//
	InputConfig interface{} `field:"optional" json:"inputConfig" yaml:"inputConfig"`
	// What type of information is found in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-dataset.html#cfn-cleanroomsml-trainingdataset-dataset-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


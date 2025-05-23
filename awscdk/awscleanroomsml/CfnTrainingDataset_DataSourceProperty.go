package awscleanroomsml


// Defines information about the Glue data source that contains the training data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceProperty := &DataSourceProperty{
//   	GlueDataSource: &GlueDataSourceProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		CatalogId: jsii.String("catalogId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasource.html
//
type CfnTrainingDataset_DataSourceProperty struct {
	// A GlueDataSource object that defines the catalog ID, database name, and table name for the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-datasource.html#cfn-cleanroomsml-trainingdataset-datasource-gluedatasource
	//
	GlueDataSource interface{} `field:"required" json:"glueDataSource" yaml:"glueDataSource"`
}


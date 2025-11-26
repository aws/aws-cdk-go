package previewawscleanroomsmlmixins


// Defines the Glue data source that contains the training data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueDataSourceProperty := &GlueDataSourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-gluedatasource.html
//
type CfnTrainingDatasetPropsMixin_GlueDataSourceProperty struct {
	// The Glue catalog that contains the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-gluedatasource.html#cfn-cleanroomsml-trainingdataset-gluedatasource-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The Glue database that contains the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-gluedatasource.html#cfn-cleanroomsml-trainingdataset-gluedatasource-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The Glue table that contains the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-gluedatasource.html#cfn-cleanroomsml-trainingdataset-gluedatasource-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}


package awscleanroomsml


// Metadata for a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSchemaProperty := &ColumnSchemaProperty{
//   	ColumnName: jsii.String("columnName"),
//   	ColumnTypes: []*string{
//   		jsii.String("columnTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-columnschema.html
//
type CfnTrainingDataset_ColumnSchemaProperty struct {
	// The name of a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-columnschema.html#cfn-cleanroomsml-trainingdataset-columnschema-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The data type of column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-trainingdataset-columnschema.html#cfn-cleanroomsml-trainingdataset-columnschema-columntypes
	//
	ColumnTypes *[]*string `field:"required" json:"columnTypes" yaml:"columnTypes"`
}


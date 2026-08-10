package awsscn


// The schema of the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   schemaProperty := &SchemaProperty{
//   	Fields: []interface{}{
//   		&DataLakeDatasetSchemaFieldProperty{
//   			IsRequired: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PrimaryKeys: []interface{}{
//   		&DataLakeDatasetPrimaryKeyFieldProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-schema.html
//
type CfnDatasetPropsMixin_SchemaProperty struct {
	// The list of field details of the dataset schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-schema.html#cfn-scn-dataset-schema-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the dataset schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-schema.html#cfn-scn-dataset-schema-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of primary key fields for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-schema.html#cfn-scn-dataset-schema-primarykeys
	//
	PrimaryKeys interface{} `field:"optional" json:"primaryKeys" yaml:"primaryKeys"`
}


package awsscn


// The dataset field details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataLakeDatasetSchemaFieldProperty := &DataLakeDatasetSchemaFieldProperty{
//   	IsRequired: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetschemafield.html
//
type CfnDatasetPropsMixin_DataLakeDatasetSchemaFieldProperty struct {
	// Indicate if the field is required or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetschemafield.html#cfn-scn-dataset-datalakedatasetschemafield-isrequired
	//
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// The dataset field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetschemafield.html#cfn-scn-dataset-datalakedatasetschemafield-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The dataset field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetschemafield.html#cfn-scn-dataset-datalakedatasetschemafield-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


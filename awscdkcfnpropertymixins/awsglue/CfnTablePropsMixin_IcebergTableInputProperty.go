package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var properties interface{}
//
//   icebergTableInputProperty := &IcebergTableInputProperty{
//   	Location: jsii.String("location"),
//   	PartitionSpec: &IcebergPartitionSpecProperty{
//   		Fields: []interface{}{
//   			&IcebergPartitionFieldProperty{
//   				FieldId: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				SourceId: jsii.Number(123),
//   				Transform: jsii.String("transform"),
//   			},
//   		},
//   		SpecId: jsii.Number(123),
//   	},
//   	Properties: properties,
//   	Schema: &IcebergSchemaProperty{
//   		Fields: []interface{}{
//   			&IcebergStructFieldProperty{
//   				Doc: jsii.String("doc"),
//   				Id: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				Required: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		IdentifierFieldIds: []interface{}{
//   			jsii.Number(123),
//   		},
//   		SchemaId: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	WriteOrder: &IcebergSortOrderProperty{
//   		Fields: []interface{}{
//   			&IcebergSortFieldProperty{
//   				Direction: jsii.String("direction"),
//   				NullOrder: jsii.String("nullOrder"),
//   				SourceId: jsii.Number(123),
//   				Transform: jsii.String("transform"),
//   			},
//   		},
//   		OrderId: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html
//
type CfnTablePropsMixin_IcebergTableInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-writeorder
	//
	WriteOrder interface{} `field:"optional" json:"writeOrder" yaml:"writeOrder"`
}


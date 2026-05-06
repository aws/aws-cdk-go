package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var properties interface{}
//
//   icebergTableInputProperty := &IcebergTableInputProperty{
//   	Location: jsii.String("location"),
//   	Schema: &IcebergSchemaProperty{
//   		Fields: []interface{}{
//   			&IcebergStructFieldProperty{
//   				Id: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				Required: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Doc: jsii.String("doc"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		IdentifierFieldIds: []interface{}{
//   			jsii.Number(123),
//   		},
//   		SchemaId: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//
//   	// the properties below are optional
//   	PartitionSpec: &IcebergPartitionSpecProperty{
//   		Fields: []interface{}{
//   			&IcebergPartitionFieldProperty{
//   				Name: jsii.String("name"),
//   				SourceId: jsii.Number(123),
//   				Transform: jsii.String("transform"),
//
//   				// the properties below are optional
//   				FieldId: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SpecId: jsii.Number(123),
//   	},
//   	Properties: properties,
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
type CfnTable_IcebergTableInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-schema
	//
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-icebergtableinput.html#cfn-glue-table-icebergtableinput-writeorder
	//
	WriteOrder interface{} `field:"optional" json:"writeOrder" yaml:"writeOrder"`
}


package mixinsawstimestream


// A Schema specifies the expected data model of the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaProperty := &SchemaProperty{
//   	CompositePartitionKey: []interface{}{
//   		&PartitionKeyProperty{
//   			EnforcementInRecord: jsii.String("enforcementInRecord"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-schema.html
//
type CfnTablePropsMixin_SchemaProperty struct {
	// A non-empty list of partition keys defining the attributes used to partition the table data.
	//
	// The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-schema.html#cfn-timestream-table-schema-compositepartitionkey
	//
	CompositePartitionKey interface{} `field:"optional" json:"compositePartitionKey" yaml:"compositePartitionKey"`
}


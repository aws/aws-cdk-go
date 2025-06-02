package awsglue


// A structure that contains schema identity fields.
//
// Either this or the `SchemaVersionId` has to be
// provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaIdProperty := &SchemaIdProperty{
//   	RegistryName: jsii.String("registryName"),
//   	SchemaArn: jsii.String("schemaArn"),
//   	SchemaName: jsii.String("schemaName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-schemaid.html
//
type CfnPartition_SchemaIdProperty struct {
	// The name of the schema registry that contains the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-schemaid.html#cfn-glue-partition-schemaid-registryname
	//
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be
	// provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-schemaid.html#cfn-glue-partition-schemaid-schemaarn
	//
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-schemaid.html#cfn-glue-partition-schemaid-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}


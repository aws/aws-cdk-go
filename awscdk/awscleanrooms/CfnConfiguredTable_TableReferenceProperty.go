package awscleanrooms


// A pointer to the dataset that underlies this table.
//
// Currently, this can only be an table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReferenceProperty := &TableReferenceProperty{
//   	Glue: &GlueTableReferenceProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-tablereference.html
//
type CfnConfiguredTable_TableReferenceProperty struct {
	// If present, a reference to the AWS Glue table referred to by this table reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-tablereference.html#cfn-cleanrooms-configuredtable-tablereference-glue
	//
	Glue interface{} `field:"required" json:"glue" yaml:"glue"`
}


package awscleanrooms


// A pointer to the dataset that underlies this table.
//
// Currently, this can only be an AWS Glue table.
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
type CfnConfiguredTable_TableReferenceProperty struct {
	// If present, a reference to the AWS Glue table referred to by this table reference.
	Glue interface{} `field:"required" json:"glue" yaml:"glue"`
}


package awscleanrooms


// A reference to a table within an AWS Glue data catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueTableReferenceProperty := &GlueTableReferenceProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   }
//
type CfnConfiguredTable_GlueTableReferenceProperty struct {
	// The name of the database the AWS Glue table belongs to.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the AWS Glue table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


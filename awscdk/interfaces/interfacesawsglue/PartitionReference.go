package interfacesawsglue


// A reference to a Partition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionReference := &PartitionReference{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	IdentifierPartitionInputValues: jsii.String("identifierPartitionInputValues"),
//   	TableName: jsii.String("tableName"),
//   }
//
type PartitionReference struct {
	// The CatalogId of the Partition resource.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The DatabaseName of the Partition resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The IdentifierPartitionInputValues of the Partition resource.
	IdentifierPartitionInputValues *string `field:"required" json:"identifierPartitionInputValues" yaml:"identifierPartitionInputValues"`
	// The TableName of the Partition resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


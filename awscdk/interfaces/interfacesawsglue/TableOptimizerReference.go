package interfacesawsglue


// A reference to a TableOptimizer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableOptimizerReference := &TableOptimizerReference{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   	Type: jsii.String("type"),
//   }
//
type TableOptimizerReference struct {
	// The CatalogId of the TableOptimizer resource.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The DatabaseName of the TableOptimizer resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The TableName of the TableOptimizer resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The Type of the TableOptimizer resource.
	Type *string `field:"required" json:"type" yaml:"type"`
}


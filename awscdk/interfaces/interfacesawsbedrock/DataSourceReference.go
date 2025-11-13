package interfacesawsbedrock


// A reference to a DataSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceReference := &DataSourceReference{
//   	DataSourceId: jsii.String("dataSourceId"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   }
//
type DataSourceReference struct {
	// The DataSourceId of the DataSource resource.
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
	// The KnowledgeBaseId of the DataSource resource.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}


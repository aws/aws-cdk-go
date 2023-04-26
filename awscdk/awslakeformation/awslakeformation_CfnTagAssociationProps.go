package awslakeformation


// Properties for defining a `CfnTagAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnTagAssociationProps := &CfnTagAssociationProps{
//   	LfTags: []interface{}{
//   		&LFTagPairProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			TagKey: jsii.String("tagKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	Resource: &ResourceProperty{
//   		Catalog: catalog,
//   		Database: &DatabaseResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Name: jsii.String("name"),
//   		},
//   		Table: &TableResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   			TableWildcard: tableWildcard,
//   		},
//   		TableWithColumns: &TableWithColumnsResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnTagAssociationProps struct {
	// A structure containing an LF-tag key-value pair.
	LfTags interface{} `field:"required" json:"lfTags" yaml:"lfTags"`
	// UTF-8 string (valid values: `DATABASE | TABLE` ).
	//
	// The resource for which the LF-tag policy applies.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
}


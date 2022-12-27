package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   }
//
type CfnTagAssociation_TableWithColumnsResourceProperty struct {
	// `CfnTagAssociation.TableWithColumnsResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.TableWithColumnsResourceProperty.ColumnNames`.
	ColumnNames *[]*string `field:"required" json:"columnNames" yaml:"columnNames"`
	// `CfnTagAssociation.TableWithColumnsResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnTagAssociation.TableWithColumnsResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}


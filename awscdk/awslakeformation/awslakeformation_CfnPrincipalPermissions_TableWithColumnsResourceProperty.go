package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	columnWildcard: &columnWildcardProperty{
//   		excludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   }
//
type CfnPrincipalPermissions_TableWithColumnsResourceProperty struct {
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.ColumnNames`.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.ColumnWildcard`.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
}


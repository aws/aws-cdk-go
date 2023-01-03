package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tableWildcard interface{}
//
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tableWildcard: tableWildcard,
//   }
//
type CfnPrincipalPermissions_TableResourceProperty struct {
	// `CfnPrincipalPermissions.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.TableResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnPrincipalPermissions.TableResourceProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnPrincipalPermissions.TableResourceProperty.TableWildcard`.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}


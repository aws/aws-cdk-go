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
type CfnTagAssociation_TableResourceProperty struct {
	// `CfnTagAssociation.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.TableResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnTagAssociation.TableResourceProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnTagAssociation.TableResourceProperty.TableWildcard`.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}


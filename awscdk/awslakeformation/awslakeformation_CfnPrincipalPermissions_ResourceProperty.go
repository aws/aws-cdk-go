package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   resourceProperty := &resourceProperty{
//   	catalog: catalog,
//   	database: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	dataCellsFilter: &dataCellsFilterResourceProperty{
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   		tableCatalogId: jsii.String("tableCatalogId"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	dataLocation: &dataLocationResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	lfTag: &lFTagKeyResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		tagKey: jsii.String("tagKey"),
//   		tagValues: []*string{
//   			jsii.String("tagValues"),
//   		},
//   	},
//   	lfTagPolicy: &lFTagPolicyResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		expression: []interface{}{
//   			&lFTagProperty{
//   				tagKey: jsii.String("tagKey"),
//   				tagValues: []*string{
//   					jsii.String("tagValues"),
//   				},
//   			},
//   		},
//   		resourceType: jsii.String("resourceType"),
//   	},
//   	table: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   		tableWildcard: tableWildcard,
//   	},
//   	tableWithColumns: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		columnWildcard: &columnWildcardProperty{
//   			excludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   	},
//   }
//
type CfnPrincipalPermissions_ResourceProperty struct {
	// `CfnPrincipalPermissions.ResourceProperty.Catalog`.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// `CfnPrincipalPermissions.ResourceProperty.Database`.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// `CfnPrincipalPermissions.ResourceProperty.DataCellsFilter`.
	DataCellsFilter interface{} `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// `CfnPrincipalPermissions.ResourceProperty.DataLocation`.
	DataLocation interface{} `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// `CfnPrincipalPermissions.ResourceProperty.LFTag`.
	LfTag interface{} `field:"optional" json:"lfTag" yaml:"lfTag"`
	// `CfnPrincipalPermissions.ResourceProperty.LFTagPolicy`.
	LfTagPolicy interface{} `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// `CfnPrincipalPermissions.ResourceProperty.Table`.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// `CfnPrincipalPermissions.ResourceProperty.TableWithColumns`.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}


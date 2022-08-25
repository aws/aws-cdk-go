package awslakeformation


// Properties for defining a `CfnDataCellsFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//
//   cfnDataCellsFilterProps := &cfnDataCellsFilterProps{
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableCatalogId: jsii.String("tableCatalogId"),
//   	tableName: jsii.String("tableName"),
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
//   	rowFilter: &rowFilterProperty{
//   		allRowsWildcard: allRowsWildcard,
//   		filterExpression: jsii.String("filterExpression"),
//   	},
//   }
//
type CfnDataCellsFilterProps struct {
	// `AWS::LakeFormation::DataCellsFilter.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `AWS::LakeFormation::DataCellsFilter.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::LakeFormation::DataCellsFilter.TableCatalogId`.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// `AWS::LakeFormation::DataCellsFilter.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// `AWS::LakeFormation::DataCellsFilter.ColumnNames`.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// `AWS::LakeFormation::DataCellsFilter.ColumnWildcard`.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// `AWS::LakeFormation::DataCellsFilter.RowFilter`.
	RowFilter interface{} `field:"optional" json:"rowFilter" yaml:"rowFilter"`
}


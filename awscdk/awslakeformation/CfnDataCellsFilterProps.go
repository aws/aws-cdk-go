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
//   cfnDataCellsFilterProps := &CfnDataCellsFilterProps{
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   	TableCatalogId: jsii.String("tableCatalogId"),
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	ColumnWildcard: &ColumnWildcardProperty{
//   		ExcludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	RowFilter: &RowFilterProperty{
//   		AllRowsWildcard: allRowsWildcard,
//   		FilterExpression: jsii.String("filterExpression"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html
//
type CfnDataCellsFilterProps struct {
	// UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// A database in the Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// The name given by the user to the data filter cell.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Catalog id string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// The ID of the catalog to which the table belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-tablecatalogid
	//
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// A table in the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// An array of UTF-8 strings.
	//
	// A list of column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-columnnames
	//
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// A wildcard with exclusions.
	//
	// You must specify either a `ColumnNames` list or the `ColumnWildCard` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-columnwildcard
	//
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// A PartiQL predicate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datacellsfilter.html#cfn-lakeformation-datacellsfilter-rowfilter
	//
	RowFilter interface{} `field:"optional" json:"rowFilter" yaml:"rowFilter"`
}


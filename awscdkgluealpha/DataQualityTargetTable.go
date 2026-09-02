package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

// The Glue table a `DataQualityRuleset` evaluates.
//
// Example:
//   var database IDatabase
//
//   glue.NewDataQualityRuleset(this, jsii.String("MyRuleset"), &DataQualityRulesetProps{
//   	RulesetName: jsii.String("my_ruleset"),
//   	Dqdl: glue.Dqdl_FromString(jsii.String("Rules = [ RowCount > 100, IsComplete \"order_id\" ]")),
//   	TargetTable: glue.DataQualityTargetTable_FromTableName(database, jsii.String("my_table")),
//   })
//
// Experimental.
type DataQualityTargetTable interface {
	// The database name of the target table.
	// Experimental.
	DatabaseName() *string
	// The table name of the target table.
	// Experimental.
	TableName() *string
}

// The jsii proxy struct for DataQualityTargetTable
type jsiiProxy_DataQualityTargetTable struct {
	_ byte // padding
}

func (j *jsiiProxy_DataQualityTargetTable) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataQualityTargetTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}


// Target an L2 table in a database.
// Experimental.
func DataQualityTargetTable_FromTable(database interfacesawsglue.IDatabaseRef, table ITable) DataQualityTargetTable {
	_init_.Initialize()

	if err := validateDataQualityTargetTable_FromTableParameters(database, table); err != nil {
		panic(err)
	}
	var returns DataQualityTargetTable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.DataQualityTargetTable",
		"fromTable",
		[]interface{}{database, table},
		&returns,
	)

	return returns
}

// Target a table by name in a database.
//
// Use this when the table is not
// modeled as an L2 construct (e.g. it is imported or created elsewhere).
// Experimental.
func DataQualityTargetTable_FromTableName(database interfacesawsglue.IDatabaseRef, tableName *string) DataQualityTargetTable {
	_init_.Initialize()

	if err := validateDataQualityTargetTable_FromTableNameParameters(database, tableName); err != nil {
		panic(err)
	}
	var returns DataQualityTargetTable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.DataQualityTargetTable",
		"fromTableName",
		[]interface{}{database, tableName},
		&returns,
	)

	return returns
}


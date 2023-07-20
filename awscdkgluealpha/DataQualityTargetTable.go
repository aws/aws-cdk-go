package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties of a DataQualityTargetTable.
//
// Example:
//   glue.NewDataQualityRuleset(this, jsii.String("MyDataQualityRuleset"), &DataQualityRulesetProps{
//   	ClientToken: jsii.String("client_token"),
//   	Description: jsii.String("description"),
//   	RulesetName: jsii.String("ruleset_name"),
//   	RulesetDqdl: jsii.String("ruleset_dqdl"),
//   	Tags: map[string]*string{
//   		"key1": jsii.String("value1"),
//   		"key2": jsii.String("value2"),
//   	},
//   	TargetTable: glue.NewDataQualityTargetTable(jsii.String("database_name"), jsii.String("table_name")),
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


// Experimental.
func NewDataQualityTargetTable(databaseName *string, tableName *string) DataQualityTargetTable {
	_init_.Initialize()

	if err := validateNewDataQualityTargetTableParameters(databaseName, tableName); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataQualityTargetTable{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.DataQualityTargetTable",
		[]interface{}{databaseName, tableName},
		&j,
	)

	return &j
}

// Experimental.
func NewDataQualityTargetTable_Override(d DataQualityTargetTable, databaseName *string, tableName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.DataQualityTargetTable",
		[]interface{}{databaseName, tableName},
		d,
	)
}


package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties of a DataQualityTargetTable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   dataQualityTargetTable := glue_alpha.NewDataQualityTargetTable(jsii.String("databaseName"), jsii.String("tableName"))
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


package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an instance of a DynamoDB table.
type ITableV2 interface {
	ITable
	// The ID of the table.
	TableId() *string
}

// The jsii proxy for ITableV2
type jsiiProxy_ITableV2 struct {
	jsiiProxy_ITable
}

func (j *jsiiProxy_ITableV2) TableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableId",
		&returns,
	)
	return returns
}


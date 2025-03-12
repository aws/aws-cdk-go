package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for processing a group of items in a single child workflow execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var batchInput interface{}
//
//   itemBatcher := awscdk.Aws_stepfunctions.NewItemBatcher(&ItemBatcherProps{
//   	BatchInput: batchInput,
//   	MaxInputBytesPerBatch: jsii.Number(123),
//   	MaxInputBytesPerBatchPath: jsii.String("maxInputBytesPerBatchPath"),
//   	MaxItemsPerBatch: jsii.Number(123),
//   	MaxItemsPerBatchPath: jsii.String("maxItemsPerBatchPath"),
//   })
//
type ItemBatcher interface {
	// Render ItemBatcher in ASL JSON format.
	Render() interface{}
	// Validate this ItemBatcher.
	ValidateItemBatcher() *[]*string
}

// The jsii proxy struct for ItemBatcher
type jsiiProxy_ItemBatcher struct {
	_ byte // padding
}

func NewItemBatcher(props *ItemBatcherProps) ItemBatcher {
	_init_.Initialize()

	if err := validateNewItemBatcherParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ItemBatcher{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ItemBatcher",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewItemBatcher_Override(i ItemBatcher, props *ItemBatcherProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ItemBatcher",
		[]interface{}{props},
		i,
	)
}

func (i *jsiiProxy_ItemBatcher) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ItemBatcher) ValidateItemBatcher() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validateItemBatcher",
		nil, // no parameters
		&returns,
	)

	return returns
}


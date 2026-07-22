package previewawsdataexchangeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.dataexchange@SchemaChangePlannedForDataSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaChangePlannedForDataSet := awscdkmixinspreview.Events.NewSchemaChangePlannedForDataSet()
//
// Experimental.
type SchemaChangePlannedForDataSet interface {
}

// The jsii proxy struct for SchemaChangePlannedForDataSet
type jsiiProxy_SchemaChangePlannedForDataSet struct {
	_ byte // padding
}

// Experimental.
func NewSchemaChangePlannedForDataSet() SchemaChangePlannedForDataSet {
	_init_.Initialize()

	j := jsiiProxy_SchemaChangePlannedForDataSet{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.SchemaChangePlannedForDataSet",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSchemaChangePlannedForDataSet_Override(s SchemaChangePlannedForDataSet) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.SchemaChangePlannedForDataSet",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Schema Change Planned for Data Set.
// Experimental.
func SchemaChangePlannedForDataSet_EventPattern(options *SchemaChangePlannedForDataSet_SchemaChangePlannedForDataSetProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSchemaChangePlannedForDataSet_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.SchemaChangePlannedForDataSet",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


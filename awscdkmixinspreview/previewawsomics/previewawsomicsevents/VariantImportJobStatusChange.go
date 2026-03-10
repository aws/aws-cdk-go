package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@VariantImportJobStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   variantImportJobStatusChange := awscdkmixinspreview.Events.NewVariantImportJobStatusChange()
//
// Experimental.
type VariantImportJobStatusChange interface {
}

// The jsii proxy struct for VariantImportJobStatusChange
type jsiiProxy_VariantImportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewVariantImportJobStatusChange() VariantImportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_VariantImportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.VariantImportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVariantImportJobStatusChange_Override(v VariantImportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.VariantImportJobStatusChange",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for Variant Import Job Status Change.
// Experimental.
func VariantImportJobStatusChange_VariantImportJobStatusChangePattern(options *VariantImportJobStatusChange_VariantImportJobStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVariantImportJobStatusChange_VariantImportJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.VariantImportJobStatusChange",
		"variantImportJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


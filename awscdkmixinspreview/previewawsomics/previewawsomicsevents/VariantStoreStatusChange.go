package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@VariantStoreStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   variantStoreStatusChange := awscdkmixinspreview.Events.NewVariantStoreStatusChange()
//
// Experimental.
type VariantStoreStatusChange interface {
}

// The jsii proxy struct for VariantStoreStatusChange
type jsiiProxy_VariantStoreStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewVariantStoreStatusChange() VariantStoreStatusChange {
	_init_.Initialize()

	j := jsiiProxy_VariantStoreStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.VariantStoreStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVariantStoreStatusChange_Override(v VariantStoreStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.VariantStoreStatusChange",
		nil, // no parameters
		v,
	)
}

// EventBridge event pattern for Variant Store Status Change.
// Experimental.
func VariantStoreStatusChange_EventPattern(options *VariantStoreStatusChange_VariantStoreStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateVariantStoreStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.VariantStoreStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


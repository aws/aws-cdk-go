package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReferenceStoreStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceStoreStatusChange := awscdkmixinspreview.Events.NewReferenceStoreStatusChange()
//
// Experimental.
type ReferenceStoreStatusChange interface {
}

// The jsii proxy struct for ReferenceStoreStatusChange
type jsiiProxy_ReferenceStoreStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReferenceStoreStatusChange() ReferenceStoreStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReferenceStoreStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReferenceStoreStatusChange_Override(r ReferenceStoreStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Reference Store Status Change.
// Experimental.
func ReferenceStoreStatusChange_ReferenceStoreStatusChangePattern(options *ReferenceStoreStatusChange_ReferenceStoreStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReferenceStoreStatusChange_ReferenceStoreStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreStatusChange",
		"referenceStoreStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


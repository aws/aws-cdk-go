package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics"
)

// EventBridge event patterns for ReferenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var referenceStoreRef IReferenceStoreRef
//
//   referenceStoreEvents := awscdkmixinspreview.Events.ReferenceStoreEvents_FromReferenceStore(referenceStoreRef)
//
// Experimental.
type ReferenceStoreEvents interface {
	// EventBridge event pattern for ReferenceStore Reference Import Job Status Change.
	// Experimental.
	ReferenceImportJobStatusChangePattern(options *ReferenceStoreEvents_ReferenceImportJobStatusChange_ReferenceImportJobStatusChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for ReferenceStore Reference Status Change.
	// Experimental.
	ReferenceStatusChangePattern(options *ReferenceStoreEvents_ReferenceStatusChange_ReferenceStatusChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ReferenceStoreEvents
type jsiiProxy_ReferenceStoreEvents struct {
	_ byte // padding
}

// Create ReferenceStoreEvents from a ReferenceStore reference.
// Experimental.
func ReferenceStoreEvents_FromReferenceStore(referenceStoreRef interfacesawsomics.IReferenceStoreRef) ReferenceStoreEvents {
	_init_.Initialize()

	if err := validateReferenceStoreEvents_FromReferenceStoreParameters(referenceStoreRef); err != nil {
		panic(err)
	}
	var returns ReferenceStoreEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents",
		"fromReferenceStore",
		[]interface{}{referenceStoreRef},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceStoreEvents) ReferenceImportJobStatusChangePattern(options *ReferenceStoreEvents_ReferenceImportJobStatusChange_ReferenceImportJobStatusChangeProps) *awsevents.EventPattern {
	if err := r.validateReferenceImportJobStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"referenceImportJobStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReferenceStoreEvents) ReferenceStatusChangePattern(options *ReferenceStoreEvents_ReferenceStatusChange_ReferenceStatusChangeProps) *awsevents.EventPattern {
	if err := r.validateReferenceStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"referenceStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


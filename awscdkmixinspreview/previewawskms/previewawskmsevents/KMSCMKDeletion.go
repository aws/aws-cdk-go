package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.kms@KMSCMKDeletion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSCMKDeletion := awscdkmixinspreview.Events.NewKMSCMKDeletion()
//
// Experimental.
type KMSCMKDeletion interface {
}

// The jsii proxy struct for KMSCMKDeletion
type jsiiProxy_KMSCMKDeletion struct {
	_ byte // padding
}

// Experimental.
func NewKMSCMKDeletion() KMSCMKDeletion {
	_init_.Initialize()

	j := jsiiProxy_KMSCMKDeletion{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKDeletion",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKMSCMKDeletion_Override(k KMSCMKDeletion) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKDeletion",
		nil, // no parameters
		k,
	)
}

// EventBridge event pattern for KMS CMK Deletion.
// Experimental.
func KMSCMKDeletion_EventPattern(options *KMSCMKDeletion_KMSCMKDeletionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateKMSCMKDeletion_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKDeletion",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


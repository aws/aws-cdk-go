package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@S3AccessPolicyStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3AccessPolicyStatusChange := awscdkmixinspreview.Events.NewS3AccessPolicyStatusChange()
//
// Experimental.
type S3AccessPolicyStatusChange interface {
}

// The jsii proxy struct for S3AccessPolicyStatusChange
type jsiiProxy_S3AccessPolicyStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewS3AccessPolicyStatusChange() S3AccessPolicyStatusChange {
	_init_.Initialize()

	j := jsiiProxy_S3AccessPolicyStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.S3AccessPolicyStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewS3AccessPolicyStatusChange_Override(s S3AccessPolicyStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.S3AccessPolicyStatusChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for S3 Access Policy Status Change.
// Experimental.
func S3AccessPolicyStatusChange_S3AccessPolicyStatusChangePattern(options *S3AccessPolicyStatusChange_S3AccessPolicyStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateS3AccessPolicyStatusChange_S3AccessPolicyStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.S3AccessPolicyStatusChange",
		"s3AccessPolicyStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


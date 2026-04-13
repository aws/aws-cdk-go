package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileDeleteCompleted := awscdkmixinspreview.Events.NewFTPSServerFileDeleteCompleted()
//
// Experimental.
type FTPSServerFileDeleteCompleted interface {
}

// The jsii proxy struct for FTPSServerFileDeleteCompleted
type jsiiProxy_FTPSServerFileDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileDeleteCompleted() FTPSServerFileDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileDeleteCompleted_Override(f FTPSServerFileDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDeleteCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Delete Completed.
// Experimental.
func FTPSServerFileDeleteCompleted_EventPattern(options *FTPSServerFileDeleteCompleted_FTPSServerFileDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileDeleteCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDeleteCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


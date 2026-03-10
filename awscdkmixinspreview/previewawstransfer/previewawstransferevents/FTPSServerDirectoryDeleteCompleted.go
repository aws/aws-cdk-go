package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerDirectoryDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerDirectoryDeleteCompleted := awscdkmixinspreview.Events.NewFTPSServerDirectoryDeleteCompleted()
//
// Experimental.
type FTPSServerDirectoryDeleteCompleted interface {
}

// The jsii proxy struct for FTPSServerDirectoryDeleteCompleted
type jsiiProxy_FTPSServerDirectoryDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerDirectoryDeleteCompleted() FTPSServerDirectoryDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerDirectoryDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerDirectoryDeleteCompleted_Override(f FTPSServerDirectoryDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryDeleteCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server Directory Delete Completed.
// Experimental.
func FTPSServerDirectoryDeleteCompleted_FtpSServerDirectoryDeleteCompletedPattern(options *FTPSServerDirectoryDeleteCompleted_FTPSServerDirectoryDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerDirectoryDeleteCompleted_FtpSServerDirectoryDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryDeleteCompleted",
		"ftpSServerDirectoryDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


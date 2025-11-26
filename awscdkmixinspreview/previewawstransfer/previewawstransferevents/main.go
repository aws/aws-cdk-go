package previewawstransferevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents",
		reflect.TypeOf((*AgreementEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "aS2MDNSendCompletedPattern", GoMethod: "AS2MDNSendCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "aS2MDNSendFailedPattern", GoMethod: "AS2MDNSendFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "aS2PayloadReceiveCompletedPattern", GoMethod: "AS2PayloadReceiveCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "aS2PayloadReceiveFailedPattern", GoMethod: "AS2PayloadReceiveFailedPattern"},
		},
		func() interface{} {
			return &jsiiProxy_AgreementEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendCompleted",
		reflect.TypeOf((*AgreementEvents_AS2MDNSendCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AgreementEvents_AS2MDNSendCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendCompleted.AS2MDNSendCompletedProps",
		reflect.TypeOf((*AgreementEvents_AS2MDNSendCompleted_AS2MDNSendCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendCompleted.S3Attributes",
		reflect.TypeOf((*AgreementEvents_AS2MDNSendCompleted_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendFailed",
		reflect.TypeOf((*AgreementEvents_AS2MDNSendFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AgreementEvents_AS2MDNSendFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendFailed.AS2MDNSendFailedProps",
		reflect.TypeOf((*AgreementEvents_AS2MDNSendFailed_AS2MDNSendFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendFailed.S3Attributes",
		reflect.TypeOf((*AgreementEvents_AS2MDNSendFailed_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2PayloadReceiveCompleted",
		reflect.TypeOf((*AgreementEvents_AS2PayloadReceiveCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AgreementEvents_AS2PayloadReceiveCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2PayloadReceiveCompleted.AS2PayloadReceiveCompletedProps",
		reflect.TypeOf((*AgreementEvents_AS2PayloadReceiveCompleted_AS2PayloadReceiveCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2PayloadReceiveCompleted.S3Attributes",
		reflect.TypeOf((*AgreementEvents_AS2PayloadReceiveCompleted_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2PayloadReceiveFailed",
		reflect.TypeOf((*AgreementEvents_AS2PayloadReceiveFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AgreementEvents_AS2PayloadReceiveFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2PayloadReceiveFailed.AS2PayloadReceiveFailedProps",
		reflect.TypeOf((*AgreementEvents_AS2PayloadReceiveFailed_AS2PayloadReceiveFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2PayloadReceiveFailed.S3Attributes",
		reflect.TypeOf((*AgreementEvents_AS2PayloadReceiveFailed_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents",
		reflect.TypeOf((*ConnectorEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "aS2MDNReceiveCompletedPattern", GoMethod: "AS2MDNReceiveCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "aS2MDNReceiveFailedPattern", GoMethod: "AS2MDNReceiveFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "aS2PayloadSendCompletedPattern", GoMethod: "AS2PayloadSendCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "aS2PayloadSendFailedPattern", GoMethod: "AS2PayloadSendFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorDirectoryListingCompletedPattern", GoMethod: "SFTPConnectorDirectoryListingCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorDirectoryListingFailedPattern", GoMethod: "SFTPConnectorDirectoryListingFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorFileRetrieveCompletedPattern", GoMethod: "SFTPConnectorFileRetrieveCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorFileRetrieveFailedPattern", GoMethod: "SFTPConnectorFileRetrieveFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorFileSendCompletedPattern", GoMethod: "SFTPConnectorFileSendCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorFileSendFailedPattern", GoMethod: "SFTPConnectorFileSendFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorRemoteDeleteCompletedPattern", GoMethod: "SFTPConnectorRemoteDeleteCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorRemoteDeleteFailedPattern", GoMethod: "SFTPConnectorRemoteDeleteFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorRemoteMoveCompletedPattern", GoMethod: "SFTPConnectorRemoteMoveCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sFTPConnectorRemoteMoveFailedPattern", GoMethod: "SFTPConnectorRemoteMoveFailedPattern"},
		},
		func() interface{} {
			return &jsiiProxy_ConnectorEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveCompleted",
		reflect.TypeOf((*ConnectorEvents_AS2MDNReceiveCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_AS2MDNReceiveCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveCompleted.AS2MDNReceiveCompletedProps",
		reflect.TypeOf((*ConnectorEvents_AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveCompleted.S3Attributes",
		reflect.TypeOf((*ConnectorEvents_AS2MDNReceiveCompleted_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveFailed",
		reflect.TypeOf((*ConnectorEvents_AS2MDNReceiveFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_AS2MDNReceiveFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveFailed.AS2MDNReceiveFailedProps",
		reflect.TypeOf((*ConnectorEvents_AS2MDNReceiveFailed_AS2MDNReceiveFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveFailed.S3Attributes",
		reflect.TypeOf((*ConnectorEvents_AS2MDNReceiveFailed_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendCompleted",
		reflect.TypeOf((*ConnectorEvents_AS2PayloadSendCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_AS2PayloadSendCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendCompleted.AS2PayloadSendCompletedProps",
		reflect.TypeOf((*ConnectorEvents_AS2PayloadSendCompleted_AS2PayloadSendCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendCompleted.S3Attributes",
		reflect.TypeOf((*ConnectorEvents_AS2PayloadSendCompleted_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendFailed",
		reflect.TypeOf((*ConnectorEvents_AS2PayloadSendFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_AS2PayloadSendFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendFailed.AS2PayloadSendFailedProps",
		reflect.TypeOf((*ConnectorEvents_AS2PayloadSendFailed_AS2PayloadSendFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendFailed.S3Attributes",
		reflect.TypeOf((*ConnectorEvents_AS2PayloadSendFailed_S3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingCompleted",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorDirectoryListingCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorDirectoryListingCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingCompleted.OutputFileLocation",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorDirectoryListingCompleted_OutputFileLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingCompleted.SFTPConnectorDirectoryListingCompletedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingFailed",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorDirectoryListingFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorDirectoryListingFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingFailed.SFTPConnectorDirectoryListingFailedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveCompleted",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileRetrieveCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorFileRetrieveCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveCompleted.LocalFileLocation",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileRetrieveCompleted_LocalFileLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveCompleted.SFTPConnectorFileRetrieveCompletedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveFailed",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileRetrieveFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorFileRetrieveFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveFailed.LocalFileLocation",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileRetrieveFailed_LocalFileLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveFailed.SFTPConnectorFileRetrieveFailedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendCompleted",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileSendCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorFileSendCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendCompleted.LocalFileLocation",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileSendCompleted_LocalFileLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendCompleted.SFTPConnectorFileSendCompletedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendFailed",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileSendFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorFileSendFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendFailed.LocalFileLocation",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileSendFailed_LocalFileLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendFailed.SFTPConnectorFileSendFailedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteCompleted",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteDeleteCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorRemoteDeleteCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteCompleted.SFTPConnectorRemoteDeleteCompletedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteFailed",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteDeleteFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorRemoteDeleteFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteFailed.SFTPConnectorRemoteDeleteFailedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteMoveCompleted",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteMoveCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorRemoteMoveCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteMoveCompleted.SFTPConnectorRemoteMoveCompletedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteMoveFailed",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteMoveFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ConnectorEvents_SFTPConnectorRemoteMoveFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteMoveFailed.SFTPConnectorRemoteMoveFailedProps",
		reflect.TypeOf((*ConnectorEvents_SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps)(nil)).Elem(),
	)
}

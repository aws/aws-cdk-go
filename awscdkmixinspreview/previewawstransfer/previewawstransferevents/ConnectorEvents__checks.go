//go:build !no_runtime_type_checking

package previewawstransferevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstransfer"
)

func (c *jsiiProxy_ConnectorEvents) validateAs2MDNReceiveCompletedPatternParameters(options *AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateAs2MDNReceiveFailedPatternParameters(options *AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateAs2PayloadSendCompletedPatternParameters(options *AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateAs2PayloadSendFailedPatternParameters(options *AS2PayloadSendFailed_AS2PayloadSendFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorDirectoryListingCompletedPatternParameters(options *SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorDirectoryListingFailedPatternParameters(options *SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorFileRetrieveCompletedPatternParameters(options *SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorFileRetrieveFailedPatternParameters(options *SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorFileSendCompletedPatternParameters(options *SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorFileSendFailedPatternParameters(options *SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorRemoteDeleteCompletedPatternParameters(options *SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorRemoteDeleteFailedPatternParameters(options *SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorRemoteMoveCompletedPatternParameters(options *SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSftPConnectorRemoteMoveFailedPatternParameters(options *SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateConnectorEvents_FromConnectorParameters(connectorRef interfacesawstransfer.IConnectorRef) error {
	if connectorRef == nil {
		return fmt.Errorf("parameter connectorRef is required, but nil was provided")
	}

	return nil
}


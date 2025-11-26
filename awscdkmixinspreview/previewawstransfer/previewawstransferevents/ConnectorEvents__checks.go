//go:build !no_runtime_type_checking

package previewawstransferevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstransfer"
)

func (c *jsiiProxy_ConnectorEvents) validateAS2MDNReceiveCompletedPatternParameters(options *ConnectorEvents_AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateAS2MDNReceiveFailedPatternParameters(options *ConnectorEvents_AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateAS2PayloadSendCompletedPatternParameters(options *ConnectorEvents_AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateAS2PayloadSendFailedPatternParameters(options *ConnectorEvents_AS2PayloadSendFailed_AS2PayloadSendFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorDirectoryListingCompletedPatternParameters(options *ConnectorEvents_SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorDirectoryListingFailedPatternParameters(options *ConnectorEvents_SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorFileRetrieveCompletedPatternParameters(options *ConnectorEvents_SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorFileRetrieveFailedPatternParameters(options *ConnectorEvents_SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorFileSendCompletedPatternParameters(options *ConnectorEvents_SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorFileSendFailedPatternParameters(options *ConnectorEvents_SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorRemoteDeleteCompletedPatternParameters(options *ConnectorEvents_SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorRemoteDeleteFailedPatternParameters(options *ConnectorEvents_SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorRemoteMoveCompletedPatternParameters(options *ConnectorEvents_SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ConnectorEvents) validateSFTPConnectorRemoteMoveFailedPatternParameters(options *ConnectorEvents_SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) error {
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


//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiotactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose"
)

func (f *jsiiProxy_FirehosePutRecordAction) validateBindParameters(rule awsiot.ITopicRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewFirehosePutRecordActionParameters(stream awskinesisfirehose.IDeliveryStream, props *FirehosePutRecordActionProps) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


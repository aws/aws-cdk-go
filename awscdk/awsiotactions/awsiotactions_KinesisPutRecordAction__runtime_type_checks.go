//go:build !no_runtime_type_checking

package awsiotactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
)

func (k *jsiiProxy_KinesisPutRecordAction) validateBindParameters(rule awsiot.ITopicRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisPutRecordActionParameters(stream awskinesis.IStream, props *KinesisPutRecordActionProps) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


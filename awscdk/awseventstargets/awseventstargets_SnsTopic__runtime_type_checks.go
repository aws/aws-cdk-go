//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

func (s *jsiiProxy_SnsTopic) validateBindParameters(_rule awsevents.IRule) error {
	if _rule == nil {
		return fmt.Errorf("parameter _rule is required, but nil was provided")
	}

	return nil
}

func validateNewSnsTopicParameters(topic awssns.ITopic, props *SnsTopicProps) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


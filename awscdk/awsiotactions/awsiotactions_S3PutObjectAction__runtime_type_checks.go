//go:build !no_runtime_type_checking

package awsiotactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

func (s *jsiiProxy_S3PutObjectAction) validateBindParameters(rule awsiot.ITopicRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewS3PutObjectActionParameters(bucket awss3.IBucket, props *S3PutObjectActionProps) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


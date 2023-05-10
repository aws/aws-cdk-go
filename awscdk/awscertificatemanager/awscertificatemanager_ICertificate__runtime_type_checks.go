//go:build !no_runtime_type_checking

package awscertificatemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

func (i *jsiiProxy_ICertificate) validateMetricDaysToExpiryParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


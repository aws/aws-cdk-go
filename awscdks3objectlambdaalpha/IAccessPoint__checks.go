//go:build !no_runtime_type_checking

package awscdks3objectlambdaalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func (i *jsiiProxy_IAccessPoint) validateVirtualHostedUrlForObjectParameters(options *awss3.VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


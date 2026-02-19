//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_IEncryptedResourceFactory) validateForResourceParameters(resource awscdk.CfnResource) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}


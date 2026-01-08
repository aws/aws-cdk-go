//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

func (i *jsiiProxy_ISchema) validateBindParameters(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


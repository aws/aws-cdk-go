package awscdkamplifyalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A source code provider.
// Experimental.
type ISourceCodeProvider interface {
	// Binds the source code provider to an app.
	// Experimental.
	Bind(app App) *SourceCodeProviderConfig
}

// The jsii proxy for ISourceCodeProvider
type jsiiProxy_ISourceCodeProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_ISourceCodeProvider) Bind(app App) *SourceCodeProviderConfig {
	if err := i.validateBindParameters(app); err != nil {
		panic(err)
	}
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{app},
		&returns,
	)

	return returns
}


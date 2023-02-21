package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// A configuration set.
type IConfigurationSet interface {
	awscdk.IResource
	// The name of the configuration set.
	ConfigurationSetName() *string
}

// The jsii proxy for IConfigurationSet
type jsiiProxy_IConfigurationSet struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IConfigurationSet) ConfigurationSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetName",
		&returns,
	)
	return returns
}


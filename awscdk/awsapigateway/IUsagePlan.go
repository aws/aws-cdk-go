package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
)

// A UsagePlan, either managed by this CDK app, or imported.
type IUsagePlan interface {
	awscdk.IResource
	// Adds an ApiKey.
	AddApiKey(apiKey IApiKey, options *AddApiKeyOptions)
	// Id of the usage plan.
	UsagePlanId() *string
}

// The jsii proxy for IUsagePlan
type jsiiProxy_IUsagePlan struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IUsagePlan) AddApiKey(apiKey IApiKey, options *AddApiKeyOptions) {
	if err := i.validateAddApiKeyParameters(apiKey, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addApiKey",
		[]interface{}{apiKey, options},
	)
}

func (j *jsiiProxy_IUsagePlan) UsagePlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePlanId",
		&returns,
	)
	return returns
}


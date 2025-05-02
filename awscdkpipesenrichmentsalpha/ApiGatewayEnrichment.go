package awscdkpipesenrichmentsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/internal"
)

// An API Gateway enrichment for a pipe.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//   var restApi restApi
//
//
//   enrichment := enrichments.NewApiGatewayEnrichment(restApi)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSomeSource(sourceQueue),
//   	Enrichment: Enrichment,
//   	Target: NewSomeTarget(targetQueue),
//   })
//
// Experimental.
type ApiGatewayEnrichment interface {
	awscdkpipesalpha.IEnrichment
	// The ARN of the enrichment resource.
	//
	// Length Constraints: Minimum length of 0. Maximum length of 1600.
	// Experimental.
	EnrichmentArn() *string
	// Bind this enrichment to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.EnrichmentParametersConfig
	// Grant the pipes role to invoke the enrichment.
	// Experimental.
	GrantInvoke(pipeRole awsiam.IRole)
}

// The jsii proxy struct for ApiGatewayEnrichment
type jsiiProxy_ApiGatewayEnrichment struct {
	internal.Type__awscdkpipesalphaIEnrichment
}

func (j *jsiiProxy_ApiGatewayEnrichment) EnrichmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enrichmentArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGatewayEnrichment(restApi awsapigateway.IRestApi, props *ApiGatewayEnrichmentProps) ApiGatewayEnrichment {
	_init_.Initialize()

	if err := validateNewApiGatewayEnrichmentParameters(restApi, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayEnrichment{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiGatewayEnrichment",
		[]interface{}{restApi, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayEnrichment_Override(a ApiGatewayEnrichment, restApi awsapigateway.IRestApi, props *ApiGatewayEnrichmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiGatewayEnrichment",
		[]interface{}{restApi, props},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayEnrichment) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.EnrichmentParametersConfig {
	if err := a.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.EnrichmentParametersConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayEnrichment) GrantInvoke(pipeRole awsiam.IRole) {
	if err := a.validateGrantInvokeParameters(pipeRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantInvoke",
		[]interface{}{pipeRole},
	)
}


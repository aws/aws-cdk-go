package awscdkpipesenrichmentsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/internal"
)

// An API Destination enrichment for a pipe.
//
// Example:
//   var sourceQueue Queue
//   var targetQueue Queue
//
//   var apiDestination ApiDestination
//
//
//   enrichment := enrichments.NewApiDestinationEnrichment(apiDestination)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSomeSource(sourceQueue),
//   	Enrichment: Enrichment,
//   	Target: NewSomeTarget(targetQueue),
//   })
//
// Experimental.
type ApiDestinationEnrichment interface {
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

// The jsii proxy struct for ApiDestinationEnrichment
type jsiiProxy_ApiDestinationEnrichment struct {
	internal.Type__awscdkpipesalphaIEnrichment
}

func (j *jsiiProxy_ApiDestinationEnrichment) EnrichmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enrichmentArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiDestinationEnrichment(destination awsevents.IApiDestination, props *ApiDestinationEnrichmentProps) ApiDestinationEnrichment {
	_init_.Initialize()

	if err := validateNewApiDestinationEnrichmentParameters(destination, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiDestinationEnrichment{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiDestinationEnrichment",
		[]interface{}{destination, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiDestinationEnrichment_Override(a ApiDestinationEnrichment, destination awsevents.IApiDestination, props *ApiDestinationEnrichmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.ApiDestinationEnrichment",
		[]interface{}{destination, props},
		a,
	)
}

func (a *jsiiProxy_ApiDestinationEnrichment) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.EnrichmentParametersConfig {
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

func (a *jsiiProxy_ApiDestinationEnrichment) GrantInvoke(pipeRole awsiam.IRole) {
	if err := a.validateGrantInvokeParameters(pipeRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantInvoke",
		[]interface{}{pipeRole},
	)
}


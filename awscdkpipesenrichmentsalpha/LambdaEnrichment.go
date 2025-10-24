package awscdkpipesenrichmentsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/internal"
)

// A Lambda enrichment for a pipe.
//
// Example:
//   var sourceQueue Queue
//   var targetQueue Queue
//
//   var enrichmentFunction Function
//
//
//   enrichment := enrichments.NewLambdaEnrichment(enrichmentFunction)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSomeSource(sourceQueue),
//   	Enrichment: Enrichment,
//   	Target: NewSomeTarget(targetQueue),
//   })
//
// Experimental.
type LambdaEnrichment interface {
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

// The jsii proxy struct for LambdaEnrichment
type jsiiProxy_LambdaEnrichment struct {
	internal.Type__awscdkpipesalphaIEnrichment
}

func (j *jsiiProxy_LambdaEnrichment) EnrichmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enrichmentArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaEnrichment(lambda awslambda.IFunction, props *LambdaEnrichmentProps) LambdaEnrichment {
	_init_.Initialize()

	if err := validateNewLambdaEnrichmentParameters(lambda, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaEnrichment{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.LambdaEnrichment",
		[]interface{}{lambda, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaEnrichment_Override(l LambdaEnrichment, lambda awslambda.IFunction, props *LambdaEnrichmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.LambdaEnrichment",
		[]interface{}{lambda, props},
		l,
	)
}

func (l *jsiiProxy_LambdaEnrichment) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.EnrichmentParametersConfig {
	if err := l.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.EnrichmentParametersConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEnrichment) GrantInvoke(pipeRole awsiam.IRole) {
	if err := l.validateGrantInvokeParameters(pipeRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"grantInvoke",
		[]interface{}{pipeRole},
	)
}


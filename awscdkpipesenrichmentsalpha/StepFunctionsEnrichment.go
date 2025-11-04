package awscdkpipesenrichmentsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha/v2/internal"
)

// A StepFunctions enrichment for a pipe.
//
// Example:
//   var sourceQueue Queue
//   var targetQueue Queue
//
//   var enrichmentStateMachine StateMachine
//
//
//   enrichment := enrichments.NewStepFunctionsEnrichment(enrichmentStateMachine)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSomeSource(sourceQueue),
//   	Enrichment: Enrichment,
//   	Target: NewSomeTarget(targetQueue),
//   })
//
// Experimental.
type StepFunctionsEnrichment interface {
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
	GrantInvoke(grantee awsiam.IRole)
}

// The jsii proxy struct for StepFunctionsEnrichment
type jsiiProxy_StepFunctionsEnrichment struct {
	internal.Type__awscdkpipesalphaIEnrichment
}

func (j *jsiiProxy_StepFunctionsEnrichment) EnrichmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enrichmentArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionsEnrichment(stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsEnrichmentProps) StepFunctionsEnrichment {
	_init_.Initialize()

	if err := validateNewStepFunctionsEnrichmentParameters(stateMachine, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionsEnrichment{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.StepFunctionsEnrichment",
		[]interface{}{stateMachine, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionsEnrichment_Override(s StepFunctionsEnrichment, stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsEnrichmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-enrichments-alpha.StepFunctionsEnrichment",
		[]interface{}{stateMachine, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionsEnrichment) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.EnrichmentParametersConfig {
	if err := s.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.EnrichmentParametersConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionsEnrichment) GrantInvoke(grantee awsiam.IRole) {
	if err := s.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantInvoke",
		[]interface{}{grantee},
	)
}


package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Enrichment step to enhance the data from the source before sending it to the target.
// Experimental.
type IEnrichment interface {
	// Bind this enrichment to a pipe.
	// Experimental.
	Bind(pipe IPipe) *EnrichmentParametersConfig
	// Grant the pipes role to invoke the enrichment.
	// Experimental.
	GrantInvoke(grantee awsiam.IRole)
	// The ARN of the enrichment resource.
	//
	// Length Constraints: Minimum length of 0. Maximum length of 1600.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-enrichment
	//
	// Experimental.
	EnrichmentArn() *string
}

// The jsii proxy for IEnrichment
type jsiiProxy_IEnrichment struct {
	_ byte // padding
}

func (i *jsiiProxy_IEnrichment) Bind(pipe IPipe) *EnrichmentParametersConfig {
	if err := i.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *EnrichmentParametersConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnrichment) GrantInvoke(grantee awsiam.IRole) {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantInvoke",
		[]interface{}{grantee},
	)
}

func (j *jsiiProxy_IEnrichment) EnrichmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enrichmentArn",
		&returns,
	)
	return returns
}


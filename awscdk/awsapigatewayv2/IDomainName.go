package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an APIGatewayV2 DomainName.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html
//
type IDomainName interface {
	interfacesawsapigatewayv2.IDomainNameRef
	awscdk.IResource
	// The custom domain name.
	Name() *string
	// The domain name associated with the regional endpoint for this custom domain name.
	RegionalDomainName() *string
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	RegionalHostedZoneId() *string
}

// The jsii proxy for IDomainName
type jsiiProxy_IDomainName struct {
	internal.Type__interfacesawsapigatewayv2IDomainNameRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDomainName) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDomainName) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) RegionalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) DomainNameRef() *interfacesawsapigatewayv2.DomainNameReference {
	var returns *interfacesawsapigatewayv2.DomainNameReference
	_jsii_.Get(
		j,
		"domainNameRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


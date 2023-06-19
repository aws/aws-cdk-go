package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Private DNS configuration for a VPC endpoint service.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var zone hostedZone
//   var vpces vpcEndpointService
//
//
//   awscdk.NewVpcEndpointServiceDomainName(this, jsii.String("EndpointDomain"), &VpcEndpointServiceDomainNameProps{
//   	EndpointService: vpces,
//   	DomainName: jsii.String("my-stuff.aws-cdk.dev"),
//   	PublicHostedZone: zone,
//   })
//
// Experimental.
type VpcEndpointServiceDomainName interface {
	awscdk.Construct
	// The domain name associated with the private DNS configuration.
	// Experimental.
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for VpcEndpointServiceDomainName
type jsiiProxy_VpcEndpointServiceDomainName struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_VpcEndpointServiceDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpointServiceDomainName) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewVpcEndpointServiceDomainName(scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) VpcEndpointServiceDomainName {
	_init_.Initialize()

	if err := validateNewVpcEndpointServiceDomainNameParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcEndpointServiceDomainName{}

	_jsii_.Create(
		"monocdk.aws_route53.VpcEndpointServiceDomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVpcEndpointServiceDomainName_Override(v VpcEndpointServiceDomainName, scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53.VpcEndpointServiceDomainName",
		[]interface{}{scope, id, props},
		v,
	)
}

func (j *jsiiProxy_VpcEndpointServiceDomainName)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func VpcEndpointServiceDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpointServiceDomainName_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53.VpcEndpointServiceDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) OnPrepare() {
	_jsii_.InvokeVoid(
		v,
		"onPrepare",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) OnSynthesize(session constructs.ISynthesisSession) {
	if err := v.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) Prepare() {
	_jsii_.InvokeVoid(
		v,
		"prepare",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) Synthesize(session awscdk.ISynthesisSession) {
	if err := v.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		[]interface{}{session},
	)
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


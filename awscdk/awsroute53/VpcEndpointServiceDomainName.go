package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Private DNS configuration for a VPC endpoint service.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var zone publicHostedZone
//   var vpces vpcEndpointService
//
//
//   awscdk.NewVpcEndpointServiceDomainName(this, jsii.String("EndpointDomain"), &VpcEndpointServiceDomainNameProps{
//   	EndpointService: vpces,
//   	DomainName: jsii.String("my-stuff.aws-cdk.dev"),
//   	PublicHostedZone: zone,
//   })
//
type VpcEndpointServiceDomainName interface {
	constructs.Construct
	// The domain name associated with the private DNS configuration.
	DomainName() *string
	SetDomainName(val *string)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for VpcEndpointServiceDomainName
type jsiiProxy_VpcEndpointServiceDomainName struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_VpcEndpointServiceDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewVpcEndpointServiceDomainName(scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) VpcEndpointServiceDomainName {
	_init_.Initialize()

	if err := validateNewVpcEndpointServiceDomainNameParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcEndpointServiceDomainName{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.VpcEndpointServiceDomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVpcEndpointServiceDomainName_Override(v VpcEndpointServiceDomainName, scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.VpcEndpointServiceDomainName",
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

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func VpcEndpointServiceDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpointServiceDomainName_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.VpcEndpointServiceDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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


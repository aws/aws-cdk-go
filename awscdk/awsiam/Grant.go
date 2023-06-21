package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Result of a grant() operation.
//
// This class is not instantiable by consumers on purpose, so that they will be
// required to call the Grant factory functions.
//
// Example:
//   var instance instance
//   var volume volume
//
//
//   attachGrant := volume.grantAttachVolumeByResourceTag(instance.GrantPrincipal, []construct{
//   	instance,
//   })
//   detachGrant := volume.grantDetachVolumeByResourceTag(instance.GrantPrincipal, []construct{
//   	instance,
//   })
//
type Grant interface {
	constructs.IDependable
	// The statement that was added to the principal's policy.
	// Deprecated: Use `principalStatements` instead.
	PrincipalStatement() PolicyStatement
	// The statements that were added to the principal's policy.
	PrincipalStatements() *[]PolicyStatement
	// The statement that was added to the resource policy.
	// Deprecated: Use `resourceStatements` instead.
	ResourceStatement() PolicyStatement
	// The statements that were added to the principal's policy.
	ResourceStatements() *[]PolicyStatement
	// Whether the grant operation was successful.
	Success() *bool
	// Make sure this grant is applied before the given constructs are deployed.
	//
	// The same as construct.node.addDependency(grant), but slightly nicer to read.
	ApplyBefore(constructs ...constructs.IConstruct)
	// Throw an error if this grant wasn't successful.
	AssertSuccess()
	// Combine two grants into a new one.
	Combine(rhs Grant) Grant
}

// The jsii proxy struct for Grant
type jsiiProxy_Grant struct {
	internal.Type__constructsIDependable
}

func (j *jsiiProxy_Grant) PrincipalStatement() PolicyStatement {
	var returns PolicyStatement
	_jsii_.Get(
		j,
		"principalStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) PrincipalStatements() *[]PolicyStatement {
	var returns *[]PolicyStatement
	_jsii_.Get(
		j,
		"principalStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) ResourceStatement() PolicyStatement {
	var returns PolicyStatement
	_jsii_.Get(
		j,
		"resourceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) ResourceStatements() *[]PolicyStatement {
	var returns *[]PolicyStatement
	_jsii_.Get(
		j,
		"resourceStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Success() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}


// Try to grant the given permissions to the given principal.
//
// Absence of a principal leads to a warning, but failing to add
// the permissions to a present principal is not an error.
func Grant_AddToPrincipal(options *GrantOnPrincipalOptions) Grant {
	_init_.Initialize()

	if err := validateGrant_AddToPrincipalParameters(options); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Grant",
		"addToPrincipal",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Add a grant both on the principal and on the resource.
//
// As long as any principal is given, granting on the principal may fail (in
// case of a non-identity principal), but granting on the resource will
// never fail.
//
// Statement will be the resource statement.
func Grant_AddToPrincipalAndResource(options *GrantOnPrincipalAndResourceOptions) Grant {
	_init_.Initialize()

	if err := validateGrant_AddToPrincipalAndResourceParameters(options); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Grant",
		"addToPrincipalAndResource",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Grant the given permissions to the principal.
//
// The permissions will be added to the principal policy primarily, falling
// back to the resource policy if necessary. The permissions must be granted
// somewhere.
//
// - Trying to grant permissions to a principal that does not admit adding to
//   the principal policy while not providing a resource with a resource policy
//   is an error.
// - Trying to grant permissions to an absent principal (possible in the
//   case of imported resources) leads to a warning being added to the
//   resource construct.
func Grant_AddToPrincipalOrResource(options *GrantWithResourceOptions) Grant {
	_init_.Initialize()

	if err := validateGrant_AddToPrincipalOrResourceParameters(options); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Grant",
		"addToPrincipalOrResource",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns a "no-op" `Grant` object which represents a "dropped grant".
//
// This can be used for e.g. imported resources where you may not be able to modify
// the resource's policy or some underlying policy which you don't know about.
func Grant_Drop(grantee IGrantable, _intent *string) Grant {
	_init_.Initialize()

	if err := validateGrant_DropParameters(grantee, _intent); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Grant",
		"drop",
		[]interface{}{grantee, _intent},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) ApplyBefore(constructs ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range constructs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"applyBefore",
		args,
	)
}

func (g *jsiiProxy_Grant) AssertSuccess() {
	_jsii_.InvokeVoid(
		g,
		"assertSuccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grant) Combine(rhs Grant) Grant {
	if err := g.validateCombineParameters(rhs); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		g,
		"combine",
		[]interface{}{rhs},
		&returns,
	)

	return returns
}


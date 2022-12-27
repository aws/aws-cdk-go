package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A type of principal that has more control over its own representation in AssumeRolePolicyDocuments.
//
// More complex types of identity providers need more control over Role's policy documents
// than simply `{ Effect: 'Allow', Action: 'AssumeRole', Principal: <Whatever> }`.
//
// If that control is necessary, they can implement `IAssumeRolePrincipal` to get full
// access to a Role's AssumeRolePolicyDocument.
type IAssumeRolePrincipal interface {
	IPrincipal
	// Add the princpial to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	AddToAssumeRolePolicy(document PolicyDocument)
}

// The jsii proxy for IAssumeRolePrincipal
type jsiiProxy_IAssumeRolePrincipal struct {
	jsiiProxy_IPrincipal
}

func (i *jsiiProxy_IAssumeRolePrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := i.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}


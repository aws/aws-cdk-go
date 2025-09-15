package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoleAlias.
// Experimental.
type IRoleAliasRef interface {
	constructs.IConstruct
	// A reference to a RoleAlias resource.
	// Experimental.
	RoleAliasRef() *RoleAliasReference
}

// The jsii proxy for IRoleAliasRef
type jsiiProxy_IRoleAliasRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRoleAliasRef) RoleAliasRef() *RoleAliasReference {
	var returns *RoleAliasReference
	_jsii_.Get(
		j,
		"roleAliasRef",
		&returns,
	)
	return returns
}


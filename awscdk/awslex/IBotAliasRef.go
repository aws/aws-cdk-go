package awslex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BotAlias.
// Experimental.
type IBotAliasRef interface {
	constructs.IConstruct
	// A reference to a BotAlias resource.
	// Experimental.
	BotAliasRef() *BotAliasReference
}

// The jsii proxy for IBotAliasRef
type jsiiProxy_IBotAliasRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBotAliasRef) BotAliasRef() *BotAliasReference {
	var returns *BotAliasReference
	_jsii_.Get(
		j,
		"botAliasRef",
		&returns,
	)
	return returns
}


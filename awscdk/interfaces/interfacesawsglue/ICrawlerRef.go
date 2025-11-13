package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Crawler.
// Experimental.
type ICrawlerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Crawler resource.
	// Experimental.
	CrawlerRef() *CrawlerReference
}

// The jsii proxy for ICrawlerRef
type jsiiProxy_ICrawlerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICrawlerRef) CrawlerRef() *CrawlerReference {
	var returns *CrawlerReference
	_jsii_.Get(
		j,
		"crawlerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICrawlerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICrawlerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


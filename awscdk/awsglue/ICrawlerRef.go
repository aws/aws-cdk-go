package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Crawler.
// Experimental.
type ICrawlerRef interface {
	constructs.IConstruct
	// A reference to a Crawler resource.
	// Experimental.
	CrawlerRef() *CrawlerReference
}

// The jsii proxy for ICrawlerRef
type jsiiProxy_ICrawlerRef struct {
	internal.Type__constructsIConstruct
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


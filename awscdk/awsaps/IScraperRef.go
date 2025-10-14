package awsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsaps/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Scraper.
// Experimental.
type IScraperRef interface {
	constructs.IConstruct
	// A reference to a Scraper resource.
	// Experimental.
	ScraperRef() *ScraperReference
}

// The jsii proxy for IScraperRef
type jsiiProxy_IScraperRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScraperRef) ScraperRef() *ScraperReference {
	var returns *ScraperReference
	_jsii_.Get(
		j,
		"scraperRef",
		&returns,
	)
	return returns
}


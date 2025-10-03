package awsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsaps/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Scraper.
// Experimental.
type IScraperRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScraperRef
type jsiiProxy_IScraperRef struct {
	internal.Type__constructsIConstruct
}


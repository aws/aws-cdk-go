package awselasticbeanstalk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticbeanstalk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationTemplate.
// Experimental.
type IConfigurationTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationTemplateRef
type jsiiProxy_IConfigurationTemplateRef struct {
	internal.Type__constructsIConstruct
}


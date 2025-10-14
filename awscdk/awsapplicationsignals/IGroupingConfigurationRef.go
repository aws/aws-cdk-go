package awsapplicationsignals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupingConfiguration.
// Experimental.
type IGroupingConfigurationRef interface {
	constructs.IConstruct
	// A reference to a GroupingConfiguration resource.
	// Experimental.
	GroupingConfigurationRef() *GroupingConfigurationReference
}

// The jsii proxy for IGroupingConfigurationRef
type jsiiProxy_IGroupingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGroupingConfigurationRef) GroupingConfigurationRef() *GroupingConfigurationReference {
	var returns *GroupingConfigurationReference
	_jsii_.Get(
		j,
		"groupingConfigurationRef",
		&returns,
	)
	return returns
}


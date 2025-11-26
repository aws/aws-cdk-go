package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworks"
)

// EventBridge event patterns for Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stackRef IStackRef
//
//   stackEvents := awscdkmixinspreview.Events.StackEvents_FromStack(stackRef)
//
// Experimental.
type StackEvents interface {
	// EventBridge event pattern for Stack OpsWorks Deployment State Change.
	// Experimental.
	OpsWorksDeploymentStateChangePattern(options *StackEvents_OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for StackEvents
type jsiiProxy_StackEvents struct {
	_ byte // padding
}

// Create StackEvents from a Stack reference.
// Experimental.
func StackEvents_FromStack(stackRef interfacesawsopsworks.IStackRef) StackEvents {
	_init_.Initialize()

	if err := validateStackEvents_FromStackParameters(stackRef); err != nil {
		panic(err)
	}
	var returns StackEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.events.StackEvents",
		"fromStack",
		[]interface{}{stackRef},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackEvents) OpsWorksDeploymentStateChangePattern(options *StackEvents_OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps) *awsevents.EventPattern {
	if err := s.validateOpsWorksDeploymentStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		s,
		"opsWorksDeploymentStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


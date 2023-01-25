package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Initialization props for the `NestedStack` construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var topic topic
//
//   nestedStackProps := &nestedStackProps{
//   	notifications: []iTopic{
//   		topic,
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	timeout: duration,
//   }
//
// Deprecated: use core.NestedStackProps instead
type NestedStackProps struct {
	// The Simple Notification Service (SNS) topics to publish stack related events.
	// Deprecated: use core.NestedStackProps instead
	Notifications *[]awssns.ITopic `field:"optional" json:"notifications" yaml:"notifications"`
	// The set value pairs that represent the parameters passed to CloudFormation when this nested stack is created.
	//
	// Each parameter has a name corresponding
	// to a parameter defined in the embedded template and a value representing
	// the value that you want to set for the parameter.
	//
	// The nested stack construct will automatically synthesize parameters in order
	// to bind references from the parent stack(s) into the nested stack.
	// Deprecated: use core.NestedStackProps instead
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The length of time that CloudFormation waits for the nested stack to reach the CREATE_COMPLETE state.
	//
	// When CloudFormation detects that the nested stack has reached the
	// CREATE_COMPLETE state, it marks the nested stack resource as
	// CREATE_COMPLETE in the parent stack and resumes creating the parent stack.
	// If the timeout period expires before the nested stack reaches
	// CREATE_COMPLETE, CloudFormation marks the nested stack as failed and rolls
	// back both the nested stack and parent stack.
	// Deprecated: use core.NestedStackProps instead
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}


package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnListener`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerProps := &CfnListenerProps{
//   	DefaultAction: &DefaultActionProperty{
//   		FixedResponse: &FixedResponseProperty{
//   			StatusCode: jsii.Number(123),
//   		},
//   		Forward: &ForwardProperty{
//   			TargetGroups: []interface{}{
//   				&WeightedTargetGroupProperty{
//   					TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   					// the properties below are optional
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Port: jsii.Number(123),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnListenerProps struct {
	// The action for the default rule.
	//
	// Each listener has a default rule. Each rule consists of a priority, one or more actions, and one or more conditions. The default rule is the rule that's used if no other rules match. Each rule must include exactly one of the following types of actions: `forward` or `fixed-response` , and it must be the last action to be performed.
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// The listener protocol HTTP or HTTPS.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The name of the listener.
	//
	// A listener name must be unique within a service. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The listener port.
	//
	// You can specify a value from `1` to `65535` . For HTTP, the default is `80` . For HTTPS, the default is `443` .
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The ID or Amazon Resource Name (ARN) of the service.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// The tags for the listener.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


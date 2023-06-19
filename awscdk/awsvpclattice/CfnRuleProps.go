package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleProps := &CfnRuleProps{
//   	Action: &ActionProperty{
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
//   	Match: &MatchProperty{
//   		HttpMatch: &HttpMatchProperty{
//   			HeaderMatches: []interface{}{
//   				&HeaderMatchProperty{
//   					Match: &HeaderMatchTypeProperty{
//   						Contains: jsii.String("contains"),
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   					},
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					CaseSensitive: jsii.Boolean(false),
//   				},
//   			},
//   			Method: jsii.String("method"),
//   			PathMatch: &PathMatchProperty{
//   				Match: &PathMatchTypeProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   				},
//
//   				// the properties below are optional
//   				CaseSensitive: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	ListenerIdentifier: jsii.String("listenerIdentifier"),
//   	Name: jsii.String("name"),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleProps struct {
	// Describes the action for a rule.
	//
	// Each rule must include exactly one of the following types of actions: `forward` or `fixed-response` , and it must be the last action to be performed.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The rule match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// The priority assigned to the rule.
	//
	// Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The ID or Amazon Resource Name (ARN) of the listener.
	ListenerIdentifier *string `field:"optional" json:"listenerIdentifier" yaml:"listenerIdentifier"`
	// The name of the rule.
	//
	// The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID or Amazon Resource Name (ARN) of the service.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// The tags for the rule.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


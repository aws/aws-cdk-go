package mixinsawsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuleMixinProps := &CfnRuleMixinProps{
//   	Action: &ActionProperty{
//   		FixedResponse: &FixedResponseProperty{
//   			StatusCode: jsii.Number(123),
//   		},
//   		Forward: &ForwardProperty{
//   			TargetGroups: []interface{}{
//   				&WeightedTargetGroupProperty{
//   					TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	ListenerIdentifier: jsii.String("listenerIdentifier"),
//   	Match: &MatchProperty{
//   		HttpMatch: &HttpMatchProperty{
//   			HeaderMatches: []interface{}{
//   				&HeaderMatchProperty{
//   					CaseSensitive: jsii.Boolean(false),
//   					Match: &HeaderMatchTypeProperty{
//   						Contains: jsii.String("contains"),
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Method: jsii.String("method"),
//   			PathMatch: &PathMatchProperty{
//   				CaseSensitive: jsii.Boolean(false),
//   				Match: &PathMatchTypeProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html
//
type CfnRuleMixinProps struct {
	// Describes the action for a rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// The ID or ARN of the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-listeneridentifier
	//
	ListenerIdentifier *string `field:"optional" json:"listenerIdentifier" yaml:"listenerIdentifier"`
	// The rule match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// The name of the rule.
	//
	// The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The priority assigned to the rule.
	//
	// Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The ID or ARN of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-serviceidentifier
	//
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// The tags for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html#cfn-vpclattice-rule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


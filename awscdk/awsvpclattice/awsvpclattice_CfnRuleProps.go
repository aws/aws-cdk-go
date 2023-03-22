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
//   cfnRuleProps := &cfnRuleProps{
//   	action: &actionProperty{
//   		forward: &forwardProperty{
//   			targetGroups: []interface{}{
//   				&weightedTargetGroupProperty{
//   					targetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   					// the properties below are optional
//   					weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	match: &matchProperty{
//   		httpMatch: &httpMatchProperty{
//   			headerMatches: []interface{}{
//   				&headerMatchProperty{
//   					match: &headerMatchTypeProperty{
//   						contains: jsii.String("contains"),
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   					},
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					caseSensitive: jsii.Boolean(false),
//   				},
//   			},
//   			method: jsii.String("method"),
//   			pathMatch: &pathMatchProperty{
//   				match: &pathMatchTypeProperty{
//   					exact: jsii.String("exact"),
//   					prefix: jsii.String("prefix"),
//   				},
//
//   				// the properties below are optional
//   				caseSensitive: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	listenerIdentifier: jsii.String("listenerIdentifier"),
//   	name: jsii.String("name"),
//   	serviceIdentifier: jsii.String("serviceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleProps struct {
	// `AWS::VpcLattice::Rule.Action`.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// `AWS::VpcLattice::Rule.Match`.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// `AWS::VpcLattice::Rule.Priority`.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// `AWS::VpcLattice::Rule.ListenerIdentifier`.
	ListenerIdentifier *string `field:"optional" json:"listenerIdentifier" yaml:"listenerIdentifier"`
	// `AWS::VpcLattice::Rule.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::VpcLattice::Rule.ServiceIdentifier`.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// `AWS::VpcLattice::Rule.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


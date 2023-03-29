package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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


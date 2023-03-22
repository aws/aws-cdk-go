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
//   cfnListenerProps := &cfnListenerProps{
//   	defaultAction: &defaultActionProperty{
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
//   	protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   	serviceIdentifier: jsii.String("serviceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnListenerProps struct {
	// `AWS::VpcLattice::Listener.DefaultAction`.
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// `AWS::VpcLattice::Listener.Protocol`.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// `AWS::VpcLattice::Listener.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::VpcLattice::Listener.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// `AWS::VpcLattice::Listener.ServiceIdentifier`.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// `AWS::VpcLattice::Listener.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


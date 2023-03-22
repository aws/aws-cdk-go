package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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


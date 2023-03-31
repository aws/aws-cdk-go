package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetGroupProps := &cfnTargetGroupProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	config: &targetGroupConfigProperty{
//   		port: jsii.Number(123),
//   		protocol: jsii.String("protocol"),
//   		vpcIdentifier: jsii.String("vpcIdentifier"),
//
//   		// the properties below are optional
//   		healthCheck: &healthCheckConfigProperty{
//   			enabled: jsii.Boolean(false),
//   			healthCheckIntervalSeconds: jsii.Number(123),
//   			healthCheckTimeoutSeconds: jsii.Number(123),
//   			healthyThresholdCount: jsii.Number(123),
//   			matcher: &matcherProperty{
//   				httpCode: jsii.String("httpCode"),
//   			},
//   			path: jsii.String("path"),
//   			port: jsii.Number(123),
//   			protocol: jsii.String("protocol"),
//   			unhealthyThresholdCount: jsii.Number(123),
//   		},
//   		protocolVersion: jsii.String("protocolVersion"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targets: []interface{}{
//   		&targetProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			port: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnTargetGroupProps struct {
	// `AWS::VpcLattice::TargetGroup.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `AWS::VpcLattice::TargetGroup.Config`.
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// `AWS::VpcLattice::TargetGroup.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::VpcLattice::TargetGroup.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::VpcLattice::TargetGroup.Targets`.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}


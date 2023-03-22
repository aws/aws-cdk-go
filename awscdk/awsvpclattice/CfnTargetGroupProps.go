package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetGroupProps := &CfnTargetGroupProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Config: &TargetGroupConfigProperty{
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   		// the properties below are optional
//   		HealthCheck: &HealthCheckConfigProperty{
//   			Enabled: jsii.Boolean(false),
//   			HealthCheckIntervalSeconds: jsii.Number(123),
//   			HealthCheckTimeoutSeconds: jsii.Number(123),
//   			HealthyThresholdCount: jsii.Number(123),
//   			Matcher: &MatcherProperty{
//   				HttpCode: jsii.String("httpCode"),
//   			},
//   			Path: jsii.String("path"),
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   			UnhealthyThresholdCount: jsii.Number(123),
//   		},
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			Port: jsii.Number(123),
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


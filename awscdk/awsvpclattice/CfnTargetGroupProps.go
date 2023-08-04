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
//   			ProtocolVersion: jsii.String("protocolVersion"),
//   			UnhealthyThresholdCount: jsii.Number(123),
//   		},
//   		IpAddressType: jsii.String("ipAddressType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html
//
type CfnTargetGroupProps struct {
	// The type of target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html#cfn-vpclattice-targetgroup-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The target group configuration.
	//
	// If the target group type is `LAMBDA` , this parameter doesn't apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html#cfn-vpclattice-targetgroup-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// The name of the target group.
	//
	// The name must be unique within the account. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html#cfn-vpclattice-targetgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags for the target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html#cfn-vpclattice-targetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Describes a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html#cfn-vpclattice-targetgroup-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}


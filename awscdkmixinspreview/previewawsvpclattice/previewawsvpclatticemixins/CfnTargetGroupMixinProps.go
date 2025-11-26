package previewawsvpclatticemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTargetGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTargetGroupMixinProps := &CfnTargetGroupMixinProps{
//   	Config: &TargetGroupConfigProperty{
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
//   		LambdaEventStructureVersion: jsii.String("lambdaEventStructureVersion"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   		VpcIdentifier: jsii.String("vpcIdentifier"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Id: jsii.String("id"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html
//
type CfnTargetGroupMixinProps struct {
	// The target group configuration.
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
	// The type of target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-targetgroup.html#cfn-vpclattice-targetgroup-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


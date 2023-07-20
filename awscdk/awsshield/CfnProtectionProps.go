package awsshield

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProtection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var block interface{}
//   var count interface{}
//
//   cfnProtectionProps := &CfnProtectionProps{
//   	Name: jsii.String("name"),
//   	ResourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	ApplicationLayerAutomaticResponseConfiguration: &ApplicationLayerAutomaticResponseConfigurationProperty{
//   		Action: &ActionProperty{
//   			Block: block,
//   			Count: count,
//   		},
//   		Status: jsii.String("status"),
//   	},
//   	HealthCheckArns: []*string{
//   		jsii.String("healthCheckArns"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html
//
type CfnProtectionProps struct {
	// The name of the protection. For example, `My CloudFront distributions` .
	//
	// > If you change the name of an existing protection, Shield Advanced deletes the protection and replaces it with a new one. While this is happening, the protection isn't available on the AWS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html#cfn-shield-protection-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN (Amazon Resource Name) of the AWS resource that is protected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html#cfn-shield-protection-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The automatic application layer DDoS mitigation settings for the protection.
	//
	// This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html#cfn-shield-protection-applicationlayerautomaticresponseconfiguration
	//
	ApplicationLayerAutomaticResponseConfiguration interface{} `field:"optional" json:"applicationLayerAutomaticResponseConfiguration" yaml:"applicationLayerAutomaticResponseConfiguration"`
	// The ARN (Amazon Resource Name) of the health check to associate with the protection.
	//
	// Health-based detection provides improved responsiveness and accuracy in attack detection and mitigation.
	//
	// You can use this option with any resource type except for Route 53 hosted zones.
	//
	// For more information, see [Configuring health-based detection using health checks](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-advanced-health-checks.html) in the *AWS Shield Advanced Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html#cfn-shield-protection-healthcheckarns
	//
	HealthCheckArns *[]*string `field:"optional" json:"healthCheckArns" yaml:"healthCheckArns"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html#cfn-shield-protection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


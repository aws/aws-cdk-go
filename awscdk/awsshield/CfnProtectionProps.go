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
type CfnProtectionProps struct {
	// The name of the protection.
	//
	// For example, `My CloudFront distributions` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN (Amazon Resource Name) of the AWS resource that is protected.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The automatic application layer DDoS mitigation settings for the protection.
	//
	// This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
	ApplicationLayerAutomaticResponseConfiguration interface{} `field:"optional" json:"applicationLayerAutomaticResponseConfiguration" yaml:"applicationLayerAutomaticResponseConfiguration"`
	// The ARN (Amazon Resource Name) of the health check to associate with the protection.
	//
	// Health-based detection provides improved responsiveness and accuracy in attack detection and mitigation.
	//
	// You can use this option with any resource type except for Route 53 hosted zones.
	//
	// For more information, see [Configuring health-based detection using health checks](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-advanced-health-checks.html) in the *AWS Shield Advanced Developer Guide* .
	HealthCheckArns *[]*string `field:"optional" json:"healthCheckArns" yaml:"healthCheckArns"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS Shield Advanced APIs or command line interface. With AWS CloudFormation , you can only add tags to resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


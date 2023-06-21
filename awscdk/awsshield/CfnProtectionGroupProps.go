package awsshield

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProtectionGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProtectionGroupProps := &CfnProtectionGroupProps{
//   	Aggregation: jsii.String("aggregation"),
//   	Pattern: jsii.String("pattern"),
//   	ProtectionGroupId: jsii.String("protectionGroupId"),
//
//   	// the properties below are optional
//   	Members: []*string{
//   		jsii.String("members"),
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProtectionGroupProps struct {
	// Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.
	//
	// - Sum - Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.
	// - Mean - Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.
	// - Max - Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront distributions and origin resources for CloudFront distributions.
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The criteria to use to choose the protected resources for inclusion in the group.
	//
	// You can include all resources that have protections, provide a list of resource ARNs (Amazon Resource Names), or include all resources of a specified resource type.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The name of the protection group.
	//
	// You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it.
	ProtectionGroupId *string `field:"required" json:"protectionGroupId" yaml:"protectionGroupId"`
	// The ARNs (Amazon Resource Names) of the resources to include in the protection group.
	//
	// You must set this when you set `Pattern` to `ARBITRARY` and you must not set it for any other `Pattern` setting.
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
	// The resource type to include in the protection group.
	//
	// All protected resources of this type are included in the protection group. You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS Shield Advanced APIs or command line interface. With AWS CloudFormation , you can only add tags to resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


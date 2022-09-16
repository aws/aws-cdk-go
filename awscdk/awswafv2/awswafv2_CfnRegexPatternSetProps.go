package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRegexPatternSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegexPatternSetProps := &cfnRegexPatternSetProps{
//   	regularExpressionList: []*string{
//   		jsii.String("regularExpressionList"),
//   	},
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRegexPatternSetProps struct {
	// The regular expression patterns in the set.
	RegularExpressionList *[]*string `field:"required" json:"regularExpressionList" yaml:"regularExpressionList"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// A description of the set that helps with identification.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the set.
	//
	// You cannot change the name after you create the set.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRegexPatternSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegexPatternSetProps := &CfnRegexPatternSetProps{
//   	RegularExpressionList: []*string{
//   		jsii.String("regularExpressionList"),
//   	},
//   	Scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html
//
type CfnRegexPatternSetProps struct {
	// The regular expression patterns in the set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html#cfn-wafv2-regexpatternset-regularexpressionlist
	//
	RegularExpressionList *[]*string `field:"required" json:"regularExpressionList" yaml:"regularExpressionList"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an  REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html#cfn-wafv2-regexpatternset-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// A description of the set that helps with identification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html#cfn-wafv2-regexpatternset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the set.
	//
	// You cannot change the name after you create the set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html#cfn-wafv2-regexpatternset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html#cfn-wafv2-regexpatternset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


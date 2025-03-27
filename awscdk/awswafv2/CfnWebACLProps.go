package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWebACL`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html
//
type CfnWebACLProps struct {
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-defaultaction
	//
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	//
	// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-visibilityconfig
	//
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// Specifies custom configurations for the associations between the web ACL and protected resources.
	//
	// Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
	//
	// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
	//
	// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-associationconfig
	//
	AssociationConfig interface{} `field:"optional" json:"associationConfig" yaml:"associationConfig"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.
	//
	// If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-captchaconfig
	//
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings.
	//
	// If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-challengeconfig
	//
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-customresponsebodies
	//
	CustomResponseBodies interface{} `field:"optional" json:"customResponseBodies" yaml:"customResponseBodies"`
	// Specifies data protection to apply to the web request data for the web ACL.
	//
	// This is a web ACL level data protection option.
	//
	// The data protection that you configure for the web ACL alters the data that's available for any other data collection activity, including your AWS WAF logging destinations, web ACL request sampling, and Amazon Security Lake data collection and management. Your other option for data protection is in the logging configuration, which only affects logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-dataprotectionconfig
	//
	DataProtectionConfig interface{} `field:"optional" json:"dataProtectionConfig" yaml:"dataProtectionConfig"`
	// A description of the web ACL that helps with identification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the web ACL.
	//
	// You cannot change the name of a web ACL after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule statements used to identify the web requests that you want to manage.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the domains that AWS WAF should accept in a web request token.
	//
	// This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-tokendomains
	//
	TokenDomains *[]*string `field:"optional" json:"tokenDomains" yaml:"tokenDomains"`
}


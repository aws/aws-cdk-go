package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A CloudFront behavior wrapper.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var keyGroup keyGroup
//   var version version
//
//   behavior := &Behavior{
//   	AllowedMethods: awscdk.Aws_cloudfront.CloudFrontAllowedMethods_GET_HEAD,
//   	CachedMethods: awscdk.*Aws_cloudfront.CloudFrontAllowedCachedMethods_GET_HEAD,
//   	Compress: jsii.Boolean(false),
//   	DefaultTtl: cdk.Duration_Minutes(jsii.Number(30)),
//   	ForwardedValues: &ForwardedValuesProperty{
//   		QueryString: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Cookies: &CookiesProperty{
//   			Forward: jsii.String("forward"),
//
//   			// the properties below are optional
//   			WhitelistedNames: []*string{
//   				jsii.String("whitelistedNames"),
//   			},
//   		},
//   		Headers: []*string{
//   			jsii.String("headers"),
//   		},
//   		QueryStringCacheKeys: []*string{
//   			jsii.String("queryStringCacheKeys"),
//   		},
//   	},
//   	FunctionAssociations: []functionAssociation{
//   		&functionAssociation{
//   			EventType: awscdk.*Aws_cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			Function: function_,
//   		},
//   	},
//   	IsDefaultBehavior: jsii.Boolean(false),
//   	LambdaFunctionAssociations: []lambdaFunctionAssociation{
//   		&lambdaFunctionAssociation{
//   			EventType: awscdk.*Aws_cloudfront.LambdaEdgeEventType_ORIGIN_REQUEST,
//   			LambdaFunction: version,
//
//   			// the properties below are optional
//   			IncludeBody: jsii.Boolean(false),
//   		},
//   	},
//   	MaxTtl: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MinTtl: cdk.Duration_*Minutes(jsii.Number(30)),
//   	PathPattern: jsii.String("pathPattern"),
//   	TrustedKeyGroups: []iKeyGroup{
//   		keyGroup,
//   	},
//   	TrustedSigners: []*string{
//   		jsii.String("trustedSigners"),
//   	},
//   	ViewerProtocolPolicy: awscdk.*Aws_cloudfront.ViewerProtocolPolicy_HTTPS_ONLY,
//   }
//
type Behavior struct {
	// The method this CloudFront distribution responds do.
	AllowedMethods CloudFrontAllowedMethods `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Which methods are cached by CloudFront by default.
	CachedMethods CloudFrontAllowedCachedMethods `field:"optional" json:"cachedMethods" yaml:"cachedMethods"`
	// If CloudFront should automatically compress some content types.
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The default amount of time CloudFront will cache an object.
	//
	// This value applies only when your custom origin does not add HTTP headers,
	// such as Cache-Control max-age, Cache-Control s-maxage, and Expires to objects.
	DefaultTtl awscdk.Duration `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// The values CloudFront will forward to the origin when making a request.
	ForwardedValues *CfnDistribution_ForwardedValuesProperty `field:"optional" json:"forwardedValues" yaml:"forwardedValues"`
	// The CloudFront functions to invoke before serving the contents.
	FunctionAssociations *[]*FunctionAssociation `field:"optional" json:"functionAssociations" yaml:"functionAssociations"`
	// If this behavior is the default behavior for the distribution.
	//
	// You must specify exactly one default distribution per CloudFront distribution.
	// The default behavior is allowed to omit the "path" property.
	IsDefaultBehavior *bool `field:"optional" json:"isDefaultBehavior" yaml:"isDefaultBehavior"`
	// Declares associated lambda@edge functions for this distribution behaviour.
	LambdaFunctionAssociations *[]*LambdaFunctionAssociation `field:"optional" json:"lambdaFunctionAssociations" yaml:"lambdaFunctionAssociations"`
	// The max amount of time you want objects to stay in the cache before CloudFront queries your origin.
	MaxTtl awscdk.Duration `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time that you want objects to stay in the cache before CloudFront queries your origin.
	MinTtl awscdk.Duration `field:"optional" json:"minTtl" yaml:"minTtl"`
	// The path this behavior responds to.
	//
	// Required for all non-default behaviors. (The default behavior implicitly has "*" as the path pattern. )
	PathPattern *string `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	TrustedKeyGroups *[]IKeyGroup `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// Trusted signers is how CloudFront allows you to serve private content.
	//
	// The signers are the account IDs that are allowed to sign cookies/presigned URLs for this distribution.
	//
	// If you pass a non empty value, all requests for this behavior must be signed (no public access will be allowed).
	// Deprecated: - We recommend using trustedKeyGroups instead of trustedSigners.
	TrustedSigners *[]*string `field:"optional" json:"trustedSigners" yaml:"trustedSigners"`
	// The viewer policy for this behavior.
	ViewerProtocolPolicy ViewerProtocolPolicy `field:"optional" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
}


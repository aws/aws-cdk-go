package previewawscloudfrontmixins


// Properties for CfnResponseHeadersPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResponseHeadersPolicyMixinProps := &CfnResponseHeadersPolicyMixinProps{
//   	ResponseHeadersPolicyConfig: &ResponseHeadersPolicyConfigProperty{
//   		Comment: jsii.String("comment"),
//   		CorsConfig: &CorsConfigProperty{
//   			AccessControlAllowCredentials: jsii.Boolean(false),
//   			AccessControlAllowHeaders: &AccessControlAllowHeadersProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlAllowMethods: &AccessControlAllowMethodsProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlAllowOrigins: &AccessControlAllowOriginsProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlExposeHeaders: &AccessControlExposeHeadersProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlMaxAgeSec: jsii.Number(123),
//   			OriginOverride: jsii.Boolean(false),
//   		},
//   		CustomHeadersConfig: &CustomHeadersConfigProperty{
//   			Items: []interface{}{
//   				&CustomHeaderProperty{
//   					Header: jsii.String("header"),
//   					Override: jsii.Boolean(false),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		RemoveHeadersConfig: &RemoveHeadersConfigProperty{
//   			Items: []interface{}{
//   				&RemoveHeaderProperty{
//   					Header: jsii.String("header"),
//   				},
//   			},
//   		},
//   		SecurityHeadersConfig: &SecurityHeadersConfigProperty{
//   			ContentSecurityPolicy: &ContentSecurityPolicyProperty{
//   				ContentSecurityPolicy: jsii.String("contentSecurityPolicy"),
//   				Override: jsii.Boolean(false),
//   			},
//   			ContentTypeOptions: &ContentTypeOptionsProperty{
//   				Override: jsii.Boolean(false),
//   			},
//   			FrameOptions: &FrameOptionsProperty{
//   				FrameOption: jsii.String("frameOption"),
//   				Override: jsii.Boolean(false),
//   			},
//   			ReferrerPolicy: &ReferrerPolicyProperty{
//   				Override: jsii.Boolean(false),
//   				ReferrerPolicy: jsii.String("referrerPolicy"),
//   			},
//   			StrictTransportSecurity: &StrictTransportSecurityProperty{
//   				AccessControlMaxAgeSec: jsii.Number(123),
//   				IncludeSubdomains: jsii.Boolean(false),
//   				Override: jsii.Boolean(false),
//   				Preload: jsii.Boolean(false),
//   			},
//   			XssProtection: &XSSProtectionProperty{
//   				ModeBlock: jsii.Boolean(false),
//   				Override: jsii.Boolean(false),
//   				Protection: jsii.Boolean(false),
//   				ReportUri: jsii.String("reportUri"),
//   			},
//   		},
//   		ServerTimingHeadersConfig: &ServerTimingHeadersConfigProperty{
//   			Enabled: jsii.Boolean(false),
//   			SamplingRate: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-responseheaderspolicy.html
//
type CfnResponseHeadersPolicyMixinProps struct {
	// A response headers policy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-responseheaderspolicy.html#cfn-cloudfront-responseheaderspolicy-responseheaderspolicyconfig
	//
	ResponseHeadersPolicyConfig interface{} `field:"optional" json:"responseHeadersPolicyConfig" yaml:"responseHeadersPolicyConfig"`
}


package awscloudfront


// Properties for defining a `CfnResponseHeadersPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResponseHeadersPolicyProps := &CfnResponseHeadersPolicyProps{
//   	ResponseHeadersPolicyConfig: &ResponseHeadersPolicyConfigProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
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
//   			OriginOverride: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AccessControlExposeHeaders: &AccessControlExposeHeadersProperty{
//   				Items: []*string{
//   					jsii.String("items"),
//   				},
//   			},
//   			AccessControlMaxAgeSec: jsii.Number(123),
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
//   				Override: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				IncludeSubdomains: jsii.Boolean(false),
//   				Preload: jsii.Boolean(false),
//   			},
//   			XssProtection: &XSSProtectionProperty{
//   				Override: jsii.Boolean(false),
//   				Protection: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				ModeBlock: jsii.Boolean(false),
//   				ReportUri: jsii.String("reportUri"),
//   			},
//   		},
//   		ServerTimingHeadersConfig: &ServerTimingHeadersConfigProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			SamplingRate: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnResponseHeadersPolicyProps struct {
	// A response headers policy configuration.
	ResponseHeadersPolicyConfig interface{} `field:"required" json:"responseHeadersPolicyConfig" yaml:"responseHeadersPolicyConfig"`
}


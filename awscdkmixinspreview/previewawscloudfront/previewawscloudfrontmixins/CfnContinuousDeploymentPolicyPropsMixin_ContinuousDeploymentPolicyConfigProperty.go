package previewawscloudfrontmixins


// Contains the configuration for a continuous deployment policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   continuousDeploymentPolicyConfigProperty := &ContinuousDeploymentPolicyConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	SingleHeaderPolicyConfig: &SingleHeaderPolicyConfigProperty{
//   		Header: jsii.String("header"),
//   		Value: jsii.String("value"),
//   	},
//   	SingleWeightPolicyConfig: &SingleWeightPolicyConfigProperty{
//   		SessionStickinessConfig: &SessionStickinessConfigProperty{
//   			IdleTtl: jsii.Number(123),
//   			MaximumTtl: jsii.Number(123),
//   		},
//   		Weight: jsii.Number(123),
//   	},
//   	StagingDistributionDnsNames: []*string{
//   		jsii.String("stagingDistributionDnsNames"),
//   	},
//   	TrafficConfig: &TrafficConfigProperty{
//   		SingleHeaderConfig: &SingleHeaderConfigProperty{
//   			Header: jsii.String("header"),
//   			Value: jsii.String("value"),
//   		},
//   		SingleWeightConfig: &SingleWeightConfigProperty{
//   			SessionStickinessConfig: &SessionStickinessConfigProperty{
//   				IdleTtl: jsii.Number(123),
//   				MaximumTtl: jsii.Number(123),
//   			},
//   			Weight: jsii.Number(123),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html
//
type CfnContinuousDeploymentPolicyPropsMixin_ContinuousDeploymentPolicyConfigProperty struct {
	// A Boolean that indicates whether this continuous deployment policy is enabled (in effect).
	//
	// When this value is `true` , this policy is enabled and in effect. When this value is `false` , this policy is not enabled and has no effect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// This configuration determines which HTTP requests are sent to the staging distribution.
	//
	// If the HTTP request contains a header and value that matches what you specify here, the request is sent to the staging distribution. Otherwise the request is sent to the primary distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-singleheaderpolicyconfig
	//
	SingleHeaderPolicyConfig interface{} `field:"optional" json:"singleHeaderPolicyConfig" yaml:"singleHeaderPolicyConfig"`
	// This configuration determines the percentage of HTTP requests that are sent to the staging distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-singleweightpolicyconfig
	//
	SingleWeightPolicyConfig interface{} `field:"optional" json:"singleWeightPolicyConfig" yaml:"singleWeightPolicyConfig"`
	// The CloudFront domain name of the staging distribution.
	//
	// For example: `d111111abcdef8.cloudfront.net` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-stagingdistributiondnsnames
	//
	StagingDistributionDnsNames *[]*string `field:"optional" json:"stagingDistributionDnsNames" yaml:"stagingDistributionDnsNames"`
	// Contains the parameters for routing production traffic from your primary to staging distributions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-trafficconfig
	//
	TrafficConfig interface{} `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
	// The type of traffic configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


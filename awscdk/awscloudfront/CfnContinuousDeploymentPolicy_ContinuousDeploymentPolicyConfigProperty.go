package awscloudfront


// Contains the configuration for a continuous deployment policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousDeploymentPolicyConfigProperty := &ContinuousDeploymentPolicyConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	StagingDistributionDnsNames: []*string{
//   		jsii.String("stagingDistributionDnsNames"),
//   	},
//
//   	// the properties below are optional
//   	SingleHeaderPolicyConfig: &SingleHeaderPolicyConfigProperty{
//   		Header: jsii.String("header"),
//   		Value: jsii.String("value"),
//   	},
//   	SingleWeightPolicyConfig: &SingleWeightPolicyConfigProperty{
//   		Weight: jsii.Number(123),
//
//   		// the properties below are optional
//   		SessionStickinessConfig: &SessionStickinessConfigProperty{
//   			IdleTtl: jsii.Number(123),
//   			MaximumTtl: jsii.Number(123),
//   		},
//   	},
//   	TrafficConfig: &TrafficConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		SingleHeaderConfig: &SingleHeaderConfigProperty{
//   			Header: jsii.String("header"),
//   			Value: jsii.String("value"),
//   		},
//   		SingleWeightConfig: &SingleWeightConfigProperty{
//   			Weight: jsii.Number(123),
//
//   			// the properties below are optional
//   			SessionStickinessConfig: &SessionStickinessConfigProperty{
//   				IdleTtl: jsii.Number(123),
//   				MaximumTtl: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html
//
type CfnContinuousDeploymentPolicy_ContinuousDeploymentPolicyConfigProperty struct {
	// A Boolean that indicates whether this continuous deployment policy is enabled (in effect).
	//
	// When this value is `true` , this policy is enabled and in effect. When this value is `false` , this policy is not enabled and has no effect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The CloudFront domain name of the staging distribution.
	//
	// For example: `d111111abcdef8.cloudfront.net` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-stagingdistributiondnsnames
	//
	StagingDistributionDnsNames *[]*string `field:"required" json:"stagingDistributionDnsNames" yaml:"stagingDistributionDnsNames"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-singleheaderpolicyconfig
	//
	SingleHeaderPolicyConfig interface{} `field:"optional" json:"singleHeaderPolicyConfig" yaml:"singleHeaderPolicyConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-singleweightpolicyconfig
	//
	SingleWeightPolicyConfig interface{} `field:"optional" json:"singleWeightPolicyConfig" yaml:"singleWeightPolicyConfig"`
	// Contains the parameters for routing production traffic from your primary to staging distributions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-trafficconfig
	//
	TrafficConfig interface{} `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html#cfn-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


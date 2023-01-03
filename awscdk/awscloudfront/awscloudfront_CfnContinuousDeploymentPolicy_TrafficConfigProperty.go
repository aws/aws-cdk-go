package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficConfigProperty := &trafficConfigProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	singleHeaderConfig: &singleHeaderConfigProperty{
//   		header: jsii.String("header"),
//   		value: jsii.String("value"),
//   	},
//   	singleWeightConfig: &singleWeightConfigProperty{
//   		weight: jsii.Number(123),
//
//   		// the properties below are optional
//   		sessionStickinessConfig: &sessionStickinessConfigProperty{
//   			idleTtl: jsii.Number(123),
//   			maximumTtl: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnContinuousDeploymentPolicy_TrafficConfigProperty struct {
	// `CfnContinuousDeploymentPolicy.TrafficConfigProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnContinuousDeploymentPolicy.TrafficConfigProperty.SingleHeaderConfig`.
	SingleHeaderConfig interface{} `field:"optional" json:"singleHeaderConfig" yaml:"singleHeaderConfig"`
	// `CfnContinuousDeploymentPolicy.TrafficConfigProperty.SingleWeightConfig`.
	SingleWeightConfig interface{} `field:"optional" json:"singleWeightConfig" yaml:"singleWeightConfig"`
}


package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnUsagePlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanProps := &cfnUsagePlanProps{
//   	apiStages: []interface{}{
//   		&apiStageProperty{
//   			apiId: jsii.String("apiId"),
//   			stage: jsii.String("stage"),
//   			throttle: map[string]interface{}{
//   				"throttleKey": &ThrottleSettingsProperty{
//   					"burstLimit": jsii.Number(123),
//   					"rateLimit": jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	quota: &quotaSettingsProperty{
//   		limit: jsii.Number(123),
//   		offset: jsii.Number(123),
//   		period: jsii.String("period"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	throttle: &throttleSettingsProperty{
//   		burstLimit: jsii.Number(123),
//   		rateLimit: jsii.Number(123),
//   	},
//   	usagePlanName: jsii.String("usagePlanName"),
//   }
//
type CfnUsagePlanProps struct {
	// The API stages to associate with this usage plan.
	ApiStages interface{} `field:"optional" json:"apiStages" yaml:"apiStages"`
	// A description of the usage plan.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures the number of requests that users can make within a given interval.
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
	// A name for the usage plan.
	UsagePlanName *string `field:"optional" json:"usagePlanName" yaml:"usagePlanName"`
}


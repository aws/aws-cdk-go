package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// The associated API stages of a usage plan.
	ApiStages interface{} `field:"optional" json:"apiStages" yaml:"apiStages"`
	// The description of a usage plan.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The target maximum number of permitted requests per a given unit time interval.
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A map containing method level throttling information for API stage in a usage plan.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of a usage plan.
	UsagePlanName *string `field:"optional" json:"usagePlanName" yaml:"usagePlanName"`
}


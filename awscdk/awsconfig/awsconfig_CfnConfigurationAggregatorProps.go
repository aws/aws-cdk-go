package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnConfigurationAggregator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationAggregatorProps := &CfnConfigurationAggregatorProps{
//   	AccountAggregationSources: []interface{}{
//   		&AccountAggregationSourceProperty{
//   			AccountIds: []*string{
//   				jsii.String("accountIds"),
//   			},
//
//   			// the properties below are optional
//   			AllAwsRegions: jsii.Boolean(false),
//   			AwsRegions: []*string{
//   				jsii.String("awsRegions"),
//   			},
//   		},
//   	},
//   	ConfigurationAggregatorName: jsii.String("configurationAggregatorName"),
//   	OrganizationAggregationSource: &OrganizationAggregationSourceProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		AllAwsRegions: jsii.Boolean(false),
//   		AwsRegions: []*string{
//   			jsii.String("awsRegions"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConfigurationAggregatorProps struct {
	// Provides a list of source accounts and regions to be aggregated.
	AccountAggregationSources interface{} `field:"optional" json:"accountAggregationSources" yaml:"accountAggregationSources"`
	// The name of the aggregator.
	ConfigurationAggregatorName *string `field:"optional" json:"configurationAggregatorName" yaml:"configurationAggregatorName"`
	// Provides an organization and list of regions to be aggregated.
	OrganizationAggregationSource interface{} `field:"optional" json:"organizationAggregationSource" yaml:"organizationAggregationSource"`
	// An array of tag object.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


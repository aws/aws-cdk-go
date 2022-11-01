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
//   cfnConfigurationAggregatorProps := &cfnConfigurationAggregatorProps{
//   	accountAggregationSources: []interface{}{
//   		&accountAggregationSourceProperty{
//   			accountIds: []*string{
//   				jsii.String("accountIds"),
//   			},
//
//   			// the properties below are optional
//   			allAwsRegions: jsii.Boolean(false),
//   			awsRegions: []*string{
//   				jsii.String("awsRegions"),
//   			},
//   		},
//   	},
//   	configurationAggregatorName: jsii.String("configurationAggregatorName"),
//   	organizationAggregationSource: &organizationAggregationSourceProperty{
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		allAwsRegions: jsii.Boolean(false),
//   		awsRegions: []*string{
//   			jsii.String("awsRegions"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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


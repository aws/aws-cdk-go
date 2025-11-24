package mixinsawsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfigurationAggregatorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationAggregatorMixinProps := &CfnConfigurationAggregatorMixinProps{
//   	AccountAggregationSources: []interface{}{
//   		&AccountAggregationSourceProperty{
//   			AccountIds: []*string{
//   				jsii.String("accountIds"),
//   			},
//   			AllAwsRegions: jsii.Boolean(false),
//   			AwsRegions: []*string{
//   				jsii.String("awsRegions"),
//   			},
//   		},
//   	},
//   	ConfigurationAggregatorName: jsii.String("configurationAggregatorName"),
//   	OrganizationAggregationSource: &OrganizationAggregationSourceProperty{
//   		AllAwsRegions: jsii.Boolean(false),
//   		AwsRegions: []*string{
//   			jsii.String("awsRegions"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html
//
type CfnConfigurationAggregatorMixinProps struct {
	// Provides a list of source accounts and regions to be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-accountaggregationsources
	//
	AccountAggregationSources interface{} `field:"optional" json:"accountAggregationSources" yaml:"accountAggregationSources"`
	// The name of the aggregator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-configurationaggregatorname
	//
	ConfigurationAggregatorName *string `field:"optional" json:"configurationAggregatorName" yaml:"configurationAggregatorName"`
	// Provides an organization and list of regions to be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-organizationaggregationsource
	//
	OrganizationAggregationSource interface{} `field:"optional" json:"organizationAggregationSource" yaml:"organizationAggregationSource"`
	// An array of tag object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


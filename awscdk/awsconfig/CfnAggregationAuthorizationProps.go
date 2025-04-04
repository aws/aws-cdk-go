package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAggregationAuthorization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAggregationAuthorizationProps := &CfnAggregationAuthorizationProps{
//   	AuthorizedAccountId: jsii.String("authorizedAccountId"),
//   	AuthorizedAwsRegion: jsii.String("authorizedAwsRegion"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html
//
type CfnAggregationAuthorizationProps struct {
	// The 12-digit account ID of the account authorized to aggregate data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html#cfn-config-aggregationauthorization-authorizedaccountid
	//
	AuthorizedAccountId *string `field:"required" json:"authorizedAccountId" yaml:"authorizedAccountId"`
	// The region authorized to collect aggregated data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html#cfn-config-aggregationauthorization-authorizedawsregion
	//
	AuthorizedAwsRegion *string `field:"required" json:"authorizedAwsRegion" yaml:"authorizedAwsRegion"`
	// An array of tag object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html#cfn-config-aggregationauthorization-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


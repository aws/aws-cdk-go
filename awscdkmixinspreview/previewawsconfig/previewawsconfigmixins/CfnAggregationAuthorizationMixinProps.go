package previewawsconfigmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAggregationAuthorizationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAggregationAuthorizationMixinProps := &CfnAggregationAuthorizationMixinProps{
//   	AuthorizedAccountId: jsii.String("authorizedAccountId"),
//   	AuthorizedAwsRegion: jsii.String("authorizedAwsRegion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html
//
type CfnAggregationAuthorizationMixinProps struct {
	// The 12-digit account ID of the account authorized to aggregate data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html#cfn-config-aggregationauthorization-authorizedaccountid
	//
	AuthorizedAccountId *string `field:"optional" json:"authorizedAccountId" yaml:"authorizedAccountId"`
	// The region authorized to collect aggregated data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html#cfn-config-aggregationauthorization-authorizedawsregion
	//
	AuthorizedAwsRegion *string `field:"optional" json:"authorizedAwsRegion" yaml:"authorizedAwsRegion"`
	// An array of tag object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html#cfn-config-aggregationauthorization-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


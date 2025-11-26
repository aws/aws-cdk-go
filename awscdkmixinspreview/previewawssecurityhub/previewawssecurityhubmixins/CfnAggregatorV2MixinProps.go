package previewawssecurityhubmixins


// Properties for CfnAggregatorV2PropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAggregatorV2MixinProps := &CfnAggregatorV2MixinProps{
//   	LinkedRegions: []*string{
//   		jsii.String("linkedRegions"),
//   	},
//   	RegionLinkingMode: jsii.String("regionLinkingMode"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html
//
type CfnAggregatorV2MixinProps struct {
	// The list of Regions that are linked to the aggregation Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html#cfn-securityhub-aggregatorv2-linkedregions
	//
	LinkedRegions *[]*string `field:"optional" json:"linkedRegions" yaml:"linkedRegions"`
	// Determines how Regions are linked to an Aggregator V2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html#cfn-securityhub-aggregatorv2-regionlinkingmode
	//
	RegionLinkingMode *string `field:"optional" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// A list of key-value pairs to be applied to the AggregatorV2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html#cfn-securityhub-aggregatorv2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


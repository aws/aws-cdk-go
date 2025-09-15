package awssecurityhub


// Properties for defining a `CfnAggregatorV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAggregatorV2Props := &CfnAggregatorV2Props{
//   	LinkedRegions: []*string{
//   		jsii.String("linkedRegions"),
//   	},
//   	RegionLinkingMode: jsii.String("regionLinkingMode"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html
//
type CfnAggregatorV2Props struct {
	// The list of Regions that are linked to the aggregation Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html#cfn-securityhub-aggregatorv2-linkedregions
	//
	LinkedRegions *[]*string `field:"required" json:"linkedRegions" yaml:"linkedRegions"`
	// Determines how Regions are linked to an Aggregator V2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html#cfn-securityhub-aggregatorv2-regionlinkingmode
	//
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// A list of key-value pairs to be applied to the AggregatorV2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-aggregatorv2.html#cfn-securityhub-aggregatorv2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


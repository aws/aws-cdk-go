package awssecurityhub


// Properties for defining a `CfnFindingAggregator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFindingAggregatorProps := &CfnFindingAggregatorProps{
//   	RegionLinkingMode: jsii.String("regionLinkingMode"),
//
//   	// the properties below are optional
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html
//
type CfnFindingAggregatorProps struct {
	// Indicates whether to aggregate findings from all of the available Regions in the current partition.
	//
	// Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
	//
	// The selected option also determines how to use the Regions provided in the Regions list.
	//
	// In CloudFormation , the options for this property are as follows:
	//
	// - `ALL_REGIONS` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	// - `ALL_REGIONS_EXCEPT_SPECIFIED` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the `Regions` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	// - `SPECIFIED_REGIONS` - Indicates to aggregate findings only from the Regions listed in the `Regions` parameter. Security Hub does not automatically aggregate findings from new Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html#cfn-securityhub-findingaggregator-regionlinkingmode
	//
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// If `RegionLinkingMode` is `ALL_REGIONS_EXCEPT_SPECIFIED` , then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.
	//
	// If `RegionLinkingMode` is `SPECIFIED_REGIONS` , then this is a space-separated list of Regions that do aggregate findings to the aggregation Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html#cfn-securityhub-findingaggregator-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


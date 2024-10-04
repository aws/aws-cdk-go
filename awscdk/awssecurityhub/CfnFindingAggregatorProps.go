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
	// The options are as follows:
	//
	// - `ALL_REGIONS` - Aggregates findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	// - `ALL_REGIONS_EXCEPT_SPECIFIED` - Aggregates findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the `Regions` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	// - `SPECIFIED_REGIONS` - Aggregates findings only from the Regions listed in the `Regions` parameter. Security Hub does not automatically aggregate findings from new Regions.
	// - `NO_REGIONS` - Aggregates no data because no Regions are selected as linked Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html#cfn-securityhub-findingaggregator-regionlinkingmode
	//
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// If `RegionLinkingMode` is `ALL_REGIONS_EXCEPT_SPECIFIED` , then this is a space-separated list of Regions that don't replicate and send findings to the home Region.
	//
	// If `RegionLinkingMode` is `SPECIFIED_REGIONS` , then this is a space-separated list of Regions that do replicate and send findings to the home Region.
	//
	// An `InvalidInputException` error results if you populate this field while `RegionLinkingMode` is `NO_REGIONS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html#cfn-securityhub-findingaggregator-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


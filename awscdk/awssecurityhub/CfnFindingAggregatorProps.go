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
	// Indicates whether to link all Regions, all Regions except for a list of excluded Regions, or a list of included Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html#cfn-securityhub-findingaggregator-regionlinkingmode
	//
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// The list of excluded Regions or included Regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-findingaggregator.html#cfn-securityhub-findingaggregator-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


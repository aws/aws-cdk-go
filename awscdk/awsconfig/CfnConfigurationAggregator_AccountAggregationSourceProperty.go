package awsconfig


// A collection of accounts and regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountAggregationSourceProperty := &AccountAggregationSourceProperty{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//
//   	// the properties below are optional
//   	AllAwsRegions: jsii.Boolean(false),
//   	AwsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html
//
type CfnConfigurationAggregator_AccountAggregationSourceProperty struct {
	// The 12-digit account ID of the account being aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-accountids
	//
	AccountIds *[]*string `field:"required" json:"accountIds" yaml:"accountIds"`
	// If true, aggregate existing AWS Config regions and future regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-allawsregions
	//
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-awsregions
	//
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
}


package awsconfig


// A collection of accounts and regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountAggregationSourceProperty := &accountAggregationSourceProperty{
//   	accountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//
//   	// the properties below are optional
//   	allAwsRegions: jsii.Boolean(false),
//   	awsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   }
//
type CfnConfigurationAggregator_AccountAggregationSourceProperty struct {
	// The 12-digit account ID of the account being aggregated.
	AccountIds *[]*string `field:"required" json:"accountIds" yaml:"accountIds"`
	// If true, aggregate existing AWS Config regions and future regions.
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
}


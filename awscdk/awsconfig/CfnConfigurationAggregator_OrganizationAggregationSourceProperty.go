package awsconfig


// This object contains regions to set up the aggregator and an IAM role to retrieve organization details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationAggregationSourceProperty := &OrganizationAggregationSourceProperty{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AllAwsRegions: jsii.Boolean(false),
//   	AwsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html
//
type CfnConfigurationAggregator_OrganizationAggregationSourceProperty struct {
	// ARN of the IAM role used to retrieve AWS Organizations details associated with the aggregator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html#cfn-config-configurationaggregator-organizationaggregationsource-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// If true, aggregate existing AWS Config regions and future regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html#cfn-config-configurationaggregator-organizationaggregationsource-allawsregions
	//
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html#cfn-config-configurationaggregator-organizationaggregationsource-awsregions
	//
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
}


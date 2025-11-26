package previewawsconfigmixins


// This object contains regions to set up the aggregator and an IAM role to retrieve organization details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   organizationAggregationSourceProperty := &OrganizationAggregationSourceProperty{
//   	AllAwsRegions: jsii.Boolean(false),
//   	AwsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html
//
type CfnConfigurationAggregatorPropsMixin_OrganizationAggregationSourceProperty struct {
	// If true, aggregate existing AWS Config regions and future regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html#cfn-config-configurationaggregator-organizationaggregationsource-allawsregions
	//
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html#cfn-config-configurationaggregator-organizationaggregationsource-awsregions
	//
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// ARN of the IAM role used to retrieve AWS Organizations details associated with the aggregator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html#cfn-config-configurationaggregator-organizationaggregationsource-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


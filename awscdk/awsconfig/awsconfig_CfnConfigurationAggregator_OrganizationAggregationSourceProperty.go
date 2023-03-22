package awsconfig


// This object contains regions to set up the aggregator and an IAM role to retrieve organization details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationAggregationSourceProperty := &organizationAggregationSourceProperty{
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	allAwsRegions: jsii.Boolean(false),
//   	awsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   }
//
type CfnConfigurationAggregator_OrganizationAggregationSourceProperty struct {
	// ARN of the IAM role used to retrieve AWS Organizations details associated with the aggregator account.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// If true, aggregate existing AWS Config regions and future regions.
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
}


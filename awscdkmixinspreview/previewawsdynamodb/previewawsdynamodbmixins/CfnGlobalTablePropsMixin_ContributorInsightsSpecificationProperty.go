package previewawsdynamodbmixins


// Configures contributor insights settings for a replica or one of its indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contributorInsightsSpecificationProperty := &ContributorInsightsSpecificationProperty{
//   	Enabled: jsii.Boolean(false),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-contributorinsightsspecification.html
//
type CfnGlobalTablePropsMixin_ContributorInsightsSpecificationProperty struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-contributorinsightsspecification.html#cfn-dynamodb-globaltable-contributorinsightsspecification-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the CloudWatch Contributor Insights mode for a global table.
	//
	// Valid values are `ACCESSED_AND_THROTTLED_KEYS` (tracks all access and throttled events) or `THROTTLED_KEYS` (tracks only throttled events). This setting determines what type of contributor insights data is collected for the global table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-contributorinsightsspecification.html#cfn-dynamodb-globaltable-contributorinsightsspecification-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}


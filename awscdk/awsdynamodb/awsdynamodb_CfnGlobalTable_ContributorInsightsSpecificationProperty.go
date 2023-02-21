package awsdynamodb


// Configures contributor insights settings for a replica or one of its indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contributorInsightsSpecificationProperty := &contributorInsightsSpecificationProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnGlobalTable_ContributorInsightsSpecificationProperty struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


package awsdynamodb


// The settings used to enable or disable CloudWatch Contributor Insights.
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
type CfnTable_ContributorInsightsSpecificationProperty struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}


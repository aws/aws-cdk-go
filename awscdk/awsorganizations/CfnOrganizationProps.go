package awsorganizations


// Properties for defining a `CfnOrganization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationProps := &CfnOrganizationProps{
//   	FeatureSet: jsii.String("featureSet"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-organization.html
//
type CfnOrganizationProps struct {
	// Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
	//
	// - `ALL`  In addition to all the features supported by the consolidated billing feature set, the management account gains access to advanced features that give you more control over accounts in your organization. For more information, see [All features](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-all) in the *AWS Organizations User Guide* .
	// - `CONSOLIDATED_BILLING`  All member accounts have their bills consolidated to and paid by the management account. For more information, see [Consolidated billing](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-cb-only) in the *AWS Organizations User Guide* .
	//
	// > The consolidated billing feature feature set isn't available for organizations in the AWS GovCloud (US) Region.
	//
	// If you don't specify this property, the default value is `ALL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-organization.html#cfn-organizations-organization-featureset
	//
	// Default: - "ALL".
	//
	FeatureSet *string `field:"optional" json:"featureSet" yaml:"featureSet"`
}


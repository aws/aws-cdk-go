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
	// - `ALL`  In addition to all the features supported by the consolidated billing feature set, the management account gains access to advanced features that give you more control over accounts in your organization. By default or if you set the `FeatureSet` property to `ALL` , the new organization is created with all features enabled and service control policies automatically enabled in the [root](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#root) . For more information, see [All features](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-all) in the *AWS Organizations User Guide* .
	// - `CONSOLIDATED_BILLING`  All member accounts have their bills consolidated to and paid by the management account. For more information, see [Consolidated billing](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-cb-only) in the *AWS Organizations User Guide* .
	//
	// The consolidated billing feature subset isn't available for organizations in the AWS GovCloud (US) Region.
	//
	// Feature set `ALL` provides the following advanced features:
	//
	// - Apply any [policy type](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html#orgs-policy-types) to any member account in the organization.
	// - Apply [service control policies (SCPs)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps.html) to member accounts that restrict the services and actions that users (including the root user) and roles in an account can access. Using SCPs you can prevent member accounts from leaving the organization.
	// - Enable [integration with supported AWS services](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html) to let those services provide functionality across all of the accounts in your organization.
	//
	// If you don't specify this property, the default value is `ALL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-organization.html#cfn-organizations-organization-featureset
	//
	// Default: - "ALL".
	//
	FeatureSet *string `field:"optional" json:"featureSet" yaml:"featureSet"`
}


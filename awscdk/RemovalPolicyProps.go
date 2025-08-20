package awscdk


// Properties for applying a removal policy.
//
// Example:
//   var scope construct
//   var parent construct
//   var bucket cfnBucket
//
//
//   // Apply DESTROY policy to all resources in a scope
//   awscdk.RemovalPolicies_Of(*scope).Destroy()
//
//   // Apply RETAIN policy to all resources in a scope
//   awscdk.RemovalPolicies_Of(*scope).Retain()
//
//   // Apply SNAPSHOT policy to all resources in a scope
//   awscdk.RemovalPolicies_Of(*scope).Snapshot()
//
//   // Apply RETAIN_ON_UPDATE_OR_DELETE policy to all resources in a scope
//   awscdk.RemovalPolicies_Of(*scope).RetainOnUpdateOrDelete()
//
//   // Apply RETAIN policy only to specific resource types
//   awscdk.RemovalPolicies_Of(parent).Retain(&RemovalPolicyProps{
//   	ApplyToResourceTypes: []*string{
//   		jsii.String("AWS::DynamoDB::Table"),
//   		bucket.CfnResourceType,
//   		rds.CfnDBInstance_CFN_RESOURCE_TYPE_NAME(),
//   	},
//   })
//
//   // Apply SNAPSHOT policy excluding specific resource types
//   awscdk.RemovalPolicies_Of(*scope).Snapshot(&RemovalPolicyProps{
//   	ExcludeResourceTypes: []*string{
//   		jsii.String("AWS::Test::Resource"),
//   	},
//   })
//
type RemovalPolicyProps struct {
	// Apply the removal policy only to specific resource types.
	//
	// Can be a CloudFormation resource type string (e.g., 'AWS::S3::Bucket').
	// Default: - apply to all resources.
	//
	ApplyToResourceTypes *[]*string `field:"optional" json:"applyToResourceTypes" yaml:"applyToResourceTypes"`
	// Exclude specific resource types from the removal policy.
	//
	// Can be a CloudFormation resource type string (e.g., 'AWS::S3::Bucket').
	// Default: - no exclusions.
	//
	ExcludeResourceTypes *[]*string `field:"optional" json:"excludeResourceTypes" yaml:"excludeResourceTypes"`
	// The priority to use when applying this policy.
	//
	// The priority affects only the order in which aspects are applied during synthesis.
	// For RemovalPolicies, the last applied policy will override previous ones.
	//
	// NOTE: Priority does NOT determine which policy "wins" when there are conflicts.
	// The order of application determines the final policy, with later policies
	// overriding earlier ones.
	// Default: - AspectPriority.MUTATING
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}


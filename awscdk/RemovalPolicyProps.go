package awscdk


// Properties for applying a removal policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   removalPolicyProps := &RemovalPolicyProps{
//   	ApplyToResourceTypes: []*string{
//   		jsii.String("applyToResourceTypes"),
//   	},
//   	ExcludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	Priority: jsii.Number(123),
//   }
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


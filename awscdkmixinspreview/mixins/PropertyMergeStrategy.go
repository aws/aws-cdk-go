package mixins


// Strategy for handling nested properties in L1 property mixins.
//
// Example:
//   var bucket CfnBucket
//
//
//   // MERGE (default): Deep merges properties with existing values
//   awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("Enabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.PropertyMergeStrategy_MERGE,
//   }))
//
//   // OVERRIDE: Replaces existing property values
//   awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("Enabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.PropertyMergeStrategy_OVERRIDE,
//   }))
//
// Experimental.
type PropertyMergeStrategy string

const (
	// Override all properties.
	// Experimental.
	PropertyMergeStrategy_OVERRIDE PropertyMergeStrategy = "OVERRIDE"
	// Deep merge nested objects, override primitives and arrays.
	// Experimental.
	PropertyMergeStrategy_MERGE PropertyMergeStrategy = "MERGE"
)


package mixins


// Options for applying CfnProperty mixins.
//
// Example:
//   // MERGE (default): Deep merges properties with existing values
//   awscdk.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("Enabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.PropertyMergeStrategy_MERGE,
//   }))
//
//   // OVERRIDE: Replaces existing property values
//   awscdk.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("Enabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.PropertyMergeStrategy_OVERRIDE,
//   }))
//
// Experimental.
type CfnPropertyMixinOptions struct {
	// Strategy for merging nested properties.
	// Default: - PropertyMergeStrategy.MERGE
	//
	// Experimental.
	Strategy PropertyMergeStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}


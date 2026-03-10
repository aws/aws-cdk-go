package mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for applying CfnProperty mixins.
//
// Example:
//   // COMBINE (default): Deep merges properties with existing values
//   awscdk.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("Enabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdk.PropertyMergeStrategy_Combine(),
//   }))
//
//   // OVERRIDE: Replaces existing property values
//   awscdk.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("Enabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdk.PropertyMergeStrategy_Override(),
//   }))
//
// Experimental.
type CfnPropertyMixinOptions struct {
	// Strategy for merging nested properties.
	// Default: - PropertyMergeStrategy.combine()
	//
	// Experimental.
	Strategy awscdk.IMergeStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}


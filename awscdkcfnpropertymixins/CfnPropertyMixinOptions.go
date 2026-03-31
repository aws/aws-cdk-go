package awscdkcfnpropertymixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for applying CfnProperty mixins.
//
// Example:
//   overrideBucket := s3.NewCfnBucket(scope, jsii.String("OverrideBucket"))
//   overrideBucket.PublicAccessBlockConfiguration = &PublicAccessBlockConfigurationProperty{
//   	BlockPublicAcls: jsii.Boolean(true),
//   }
//
//   // Replaces the entire publicAccessBlockConfiguration
//   overrideBucket.With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
//   	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
//   		BlockPublicPolicy: jsii.Boolean(true),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdk.PropertyMergeStrategy_Override(),
//   }))
//
type CfnPropertyMixinOptions struct {
	// Strategy for merging nested properties.
	// Default: - PropertyMergeStrategy.combine()
	//
	Strategy awscdk.IMergeStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}


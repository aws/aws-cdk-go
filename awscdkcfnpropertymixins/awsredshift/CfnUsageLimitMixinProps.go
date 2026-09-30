package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnUsageLimitPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUsageLimitMixinProps := &CfnUsageLimitMixinProps{
//   	Amount: jsii.Number(123),
//   	BreachAction: jsii.String("breachAction"),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	FeatureType: jsii.String("featureType"),
//   	LimitType: jsii.String("limitType"),
//   	Period: jsii.String("period"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html
//
type CfnUsageLimitMixinProps struct {
	// The limit amount.
	//
	// If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-amount
	//
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// The action that Amazon Redshift takes when the limit is reached.
	//
	// The default is log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-breachaction
	//
	BreachAction *string `field:"optional" json:"breachAction" yaml:"breachAction"`
	// The identifier of the cluster that you want to limit usage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-clusteridentifier
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The Amazon Redshift feature that you want to limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-featuretype
	//
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
	// The type of limit.
	//
	// Depending on the feature type, this can be based on a time duration or data size. If FeatureType is spectrum, then LimitType must be data-scanned. If FeatureType is concurrency-scaling, then LimitType must be time. If FeatureType is cross-region-datasharing, then LimitType must be data-scanned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-limittype
	//
	LimitType *string `field:"optional" json:"limitType" yaml:"limitType"`
	// The time period that the amount applies to.
	//
	// A weekly period begins on Sunday. The default is monthly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-period
	//
	Period *string `field:"optional" json:"period" yaml:"period"`
	// A list of tag instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-usagelimit.html#cfn-redshift-usagelimit-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


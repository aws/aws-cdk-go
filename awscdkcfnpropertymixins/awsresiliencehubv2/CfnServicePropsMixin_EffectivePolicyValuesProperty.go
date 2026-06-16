package awsresiliencehubv2


// Effective policy values computed from the associated policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   effectivePolicyValuesProperty := &EffectivePolicyValuesProperty{
//   	AvailabilitySlo: &SloSourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.Number(123),
//   	},
//   	MultiAzDrApproach: &DisasterRecoverySourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.String("value"),
//   	},
//   	MultiAzRpo: &TargetSourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.Number(123),
//   	},
//   	MultiAzRto: &TargetSourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.Number(123),
//   	},
//   	MultiRegionDrApproach: &DisasterRecoverySourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.String("value"),
//   	},
//   	MultiRegionRpo: &TargetSourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.Number(123),
//   	},
//   	MultiRegionRto: &TargetSourceProperty{
//   		PolicyName: jsii.String("policyName"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html
//
type CfnServicePropsMixin_EffectivePolicyValuesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-availabilityslo
	//
	AvailabilitySlo interface{} `field:"optional" json:"availabilitySlo" yaml:"availabilitySlo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-multiazdrapproach
	//
	MultiAzDrApproach interface{} `field:"optional" json:"multiAzDrApproach" yaml:"multiAzDrApproach"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-multiazrpo
	//
	MultiAzRpo interface{} `field:"optional" json:"multiAzRpo" yaml:"multiAzRpo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-multiazrto
	//
	MultiAzRto interface{} `field:"optional" json:"multiAzRto" yaml:"multiAzRto"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-multiregiondrapproach
	//
	MultiRegionDrApproach interface{} `field:"optional" json:"multiRegionDrApproach" yaml:"multiRegionDrApproach"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-multiregionrpo
	//
	MultiRegionRpo interface{} `field:"optional" json:"multiRegionRpo" yaml:"multiRegionRpo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-effectivepolicyvalues.html#cfn-resiliencehubv2-service-effectivepolicyvalues-multiregionrto
	//
	MultiRegionRto interface{} `field:"optional" json:"multiRegionRto" yaml:"multiRegionRto"`
}


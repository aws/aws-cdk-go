package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   quotaSharePreemptionConfigurationProperty := &QuotaSharePreemptionConfigurationProperty{
//   	InSharePreemption: jsii.String("inSharePreemption"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotasharepreemptionconfiguration.html
//
type CfnQuotaSharePropsMixin_QuotaSharePreemptionConfigurationProperty struct {
	// Whether preemption is enabled within the quota share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotasharepreemptionconfiguration.html#cfn-batch-quotashare-quotasharepreemptionconfiguration-insharepreemption
	//
	InSharePreemption *string `field:"optional" json:"inSharePreemption" yaml:"inSharePreemption"`
}


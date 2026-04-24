package awsbatch


// Properties for CfnQuotaSharePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnQuotaShareMixinProps := &CfnQuotaShareMixinProps{
//   	CapacityLimits: []interface{}{
//   		&QuotaShareCapacityLimitProperty{
//   			CapacityUnit: jsii.String("capacityUnit"),
//   			MaxCapacity: jsii.Number(123),
//   		},
//   	},
//   	JobQueue: jsii.String("jobQueue"),
//   	PreemptionConfiguration: &QuotaSharePreemptionConfigurationProperty{
//   		InSharePreemption: jsii.String("inSharePreemption"),
//   	},
//   	QuotaShareName: jsii.String("quotaShareName"),
//   	ResourceSharingConfiguration: &QuotaShareResourceSharingConfigurationProperty{
//   		BorrowLimit: jsii.Number(123),
//   		Strategy: jsii.String("strategy"),
//   	},
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html
//
type CfnQuotaShareMixinProps struct {
	// A list that specifies the quantity and type of compute capacity allocated to the quota share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-capacitylimits
	//
	CapacityLimits interface{} `field:"optional" json:"capacityLimits" yaml:"capacityLimits"`
	// The AWS Batch job queue associated with the quota share.
	//
	// This can be the job queue name or ARN. A job queue must be in the `VALID` state before you can associate it with a quota share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-jobqueue
	//
	JobQueue *string `field:"optional" json:"jobQueue" yaml:"jobQueue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-preemptionconfiguration
	//
	PreemptionConfiguration interface{} `field:"optional" json:"preemptionConfiguration" yaml:"preemptionConfiguration"`
	// The name of the quota share.
	//
	// It can be up to 128 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-quotasharename
	//
	QuotaShareName *string `field:"optional" json:"quotaShareName" yaml:"quotaShareName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-resourcesharingconfiguration
	//
	ResourceSharingConfiguration interface{} `field:"optional" json:"resourceSharingConfiguration" yaml:"resourceSharingConfiguration"`
	// The state of the quota share.
	//
	// If the quota share is `ENABLED`, it is able to accept jobs. If the quota share is `DISABLED`, new jobs won't be accepted but jobs already submitted can finish. The default state is `ENABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags that you apply to the quota share to help you categorize and organize your resources.
	//
	// Each tag consists of a key and an optional value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-quotashare.html#cfn-batch-quotashare-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


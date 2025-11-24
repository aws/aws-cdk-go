package mixinsawsbatch


// Properties for CfnSchedulingPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchedulingPolicyMixinProps := &CfnSchedulingPolicyMixinProps{
//   	FairsharePolicy: &FairsharePolicyProperty{
//   		ComputeReservation: jsii.Number(123),
//   		ShareDecaySeconds: jsii.Number(123),
//   		ShareDistribution: []interface{}{
//   			&ShareAttributesProperty{
//   				ShareIdentifier: jsii.String("shareIdentifier"),
//   				WeightFactor: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-schedulingpolicy.html
//
type CfnSchedulingPolicyMixinProps struct {
	// The fair-share scheduling policy details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-schedulingpolicy.html#cfn-batch-schedulingpolicy-fairsharepolicy
	//
	FairsharePolicy interface{} `field:"optional" json:"fairsharePolicy" yaml:"fairsharePolicy"`
	// The name of the fair-share scheduling policy.
	//
	// It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-schedulingpolicy.html#cfn-batch-schedulingpolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags that you apply to the scheduling policy to help you categorize and organize your resources.
	//
	// Each tag consists of a key and an optional value. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in *AWS General Reference* .
	//
	// These tags can be updated or removed using the [TagResource](https://docs.aws.amazon.com/batch/latest/APIReference/API_TagResource.html) and [UntagResource](https://docs.aws.amazon.com/batch/latest/APIReference/API_UntagResource.html) API operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-schedulingpolicy.html#cfn-batch-schedulingpolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


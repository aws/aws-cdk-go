package awsbatch


// Properties for defining a `CfnSchedulingPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchedulingPolicyProps := &cfnSchedulingPolicyProps{
//   	fairsharePolicy: &fairsharePolicyProperty{
//   		computeReservation: jsii.Number(123),
//   		shareDecaySeconds: jsii.Number(123),
//   		shareDistribution: []interface{}{
//   			&shareAttributesProperty{
//   				shareIdentifier: jsii.String("shareIdentifier"),
//   				weightFactor: jsii.Number(123),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSchedulingPolicyProps struct {
	// The fair share policy of the scheduling policy.
	FairsharePolicy interface{} `field:"optional" json:"fairsharePolicy" yaml:"fairsharePolicy"`
	// The name of the scheduling policy.
	//
	// It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags that you apply to the scheduling policy to help you categorize and organize your resources.
	//
	// Each tag consists of a key and an optional value. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in *AWS General Reference* .
	//
	// These tags can be updated or removed using the [TagResource](https://docs.aws.amazon.com/batch/latest/APIReference/API_TagResource.html) and [UntagResource](https://docs.aws.amazon.com/batch/latest/APIReference/API_UntagResource.html) API operations.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


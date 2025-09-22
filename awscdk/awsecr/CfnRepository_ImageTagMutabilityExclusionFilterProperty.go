package awsecr


// Overrides the default image tag mutability setting of the repository for image tags that match the specified filters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageTagMutabilityExclusionFilterProperty := &ImageTagMutabilityExclusionFilterProperty{
//   	ImageTagMutabilityExclusionFilterType: jsii.String("imageTagMutabilityExclusionFilterType"),
//   	ImageTagMutabilityExclusionFilterValue: jsii.String("imageTagMutabilityExclusionFilterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-imagetagmutabilityexclusionfilter.html
//
type CfnRepository_ImageTagMutabilityExclusionFilterProperty struct {
	// Specifies the type of filter to use for excluding image tags from the repository's mutability setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-imagetagmutabilityexclusionfilter.html#cfn-ecr-repository-imagetagmutabilityexclusionfilter-imagetagmutabilityexclusionfiltertype
	//
	ImageTagMutabilityExclusionFilterType *string `field:"required" json:"imageTagMutabilityExclusionFilterType" yaml:"imageTagMutabilityExclusionFilterType"`
	// The value to use when filtering image tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-imagetagmutabilityexclusionfilter.html#cfn-ecr-repository-imagetagmutabilityexclusionfilter-imagetagmutabilityexclusionfiltervalue
	//
	ImageTagMutabilityExclusionFilterValue *string `field:"required" json:"imageTagMutabilityExclusionFilterValue" yaml:"imageTagMutabilityExclusionFilterValue"`
}


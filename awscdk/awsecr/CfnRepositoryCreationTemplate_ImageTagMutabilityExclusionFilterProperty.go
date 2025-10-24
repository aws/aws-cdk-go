package awsecr


// A filter that specifies which image tags should be excluded from the repository's image tag mutability setting.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilter.html
//
type CfnRepositoryCreationTemplate_ImageTagMutabilityExclusionFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilter.html#cfn-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilter-imagetagmutabilityexclusionfiltertype
	//
	ImageTagMutabilityExclusionFilterType *string `field:"required" json:"imageTagMutabilityExclusionFilterType" yaml:"imageTagMutabilityExclusionFilterType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilter.html#cfn-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilter-imagetagmutabilityexclusionfiltervalue
	//
	ImageTagMutabilityExclusionFilterValue *string `field:"required" json:"imageTagMutabilityExclusionFilterValue" yaml:"imageTagMutabilityExclusionFilterValue"`
}


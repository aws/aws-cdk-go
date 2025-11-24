package mixinsawsdeadline


// Properties for CfnLimitPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLimitMixinProps := &CfnLimitMixinProps{
//   	AmountRequirementName: jsii.String("amountRequirementName"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	MaxCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html
//
type CfnLimitMixinProps struct {
	// The value that you specify as the `name` in the `amounts` field of the `hostRequirements` in a step of a job template to declare the limit requirement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html#cfn-deadline-limit-amountrequirementname
	//
	AmountRequirementName *string `field:"optional" json:"amountRequirementName" yaml:"amountRequirementName"`
	// A description of the limit. A clear description helps you identify the purpose of the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html#cfn-deadline-limit-description
	//
	// Default: - "".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the limit used in lists to identify the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html#cfn-deadline-limit-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The unique identifier of the farm that contains the limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html#cfn-deadline-limit-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// The maximum number of resources constrained by this limit.
	//
	// When all of the resources are in use, steps that require the limit won't be scheduled until the resource is available.
	//
	// The `maxValue` must not be 0. If the value is -1, there is no restriction on the number of resources that can be acquired for this limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html#cfn-deadline-limit-maxcount
	//
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
}


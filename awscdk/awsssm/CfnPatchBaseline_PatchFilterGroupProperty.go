package awsssm


// The `PatchFilterGroup` property type specifies a set of patch filters for an AWS Systems Manager patch baseline, typically used for approval rules for a Systems Manager patch baseline.
//
// `PatchFilterGroup` is the property type for the `GlobalFilters` property of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource and the `PatchFilterGroup` property of the [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchFilterGroupProperty := &PatchFilterGroupProperty{
//   	PatchFilters: []interface{}{
//   		&PatchFilterProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfiltergroup.html
//
type CfnPatchBaseline_PatchFilterGroupProperty struct {
	// The set of patch filters that make up the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfiltergroup.html#cfn-ssm-patchbaseline-patchfiltergroup-patchfilters
	//
	PatchFilters interface{} `field:"optional" json:"patchFilters" yaml:"patchFilters"`
}


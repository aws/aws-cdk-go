package mixinsawsdlm


// *[Default policies only]* Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs.
//
// The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var excludeTags interface{}
//   var excludeVolumeTypes interface{}
//
//   exclusionsProperty := &ExclusionsProperty{
//   	ExcludeBootVolumes: jsii.Boolean(false),
//   	ExcludeTags: excludeTags,
//   	ExcludeVolumeTypes: excludeVolumeTypes,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html
//
type CfnLifecyclePolicyPropsMixin_ExclusionsProperty struct {
	// *[Default policies for EBS snapshots only]* Indicates whether to exclude volumes that are attached to instances as the boot volume.
	//
	// If you exclude boot volumes, only volumes attached as data (non-boot) volumes will be backed up by the policy. To exclude boot volumes, specify `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html#cfn-dlm-lifecyclepolicy-exclusions-excludebootvolumes
	//
	ExcludeBootVolumes interface{} `field:"optional" json:"excludeBootVolumes" yaml:"excludeBootVolumes"`
	// *[Default policies for EBS-backed AMIs only]* Specifies whether to exclude volumes that have specific tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html#cfn-dlm-lifecyclepolicy-exclusions-excludetags
	//
	ExcludeTags interface{} `field:"optional" json:"excludeTags" yaml:"excludeTags"`
	// *[Default policies for EBS snapshots only]* Specifies the volume types to exclude.
	//
	// Volumes of the specified types will not be targeted by the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-exclusions.html#cfn-dlm-lifecyclepolicy-exclusions-excludevolumetypes
	//
	ExcludeVolumeTypes interface{} `field:"optional" json:"excludeVolumeTypes" yaml:"excludeVolumeTypes"`
}


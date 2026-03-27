package awsdlm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyTargetProperty := &CrossRegionCopyTargetProperty{
//   	TargetRegion: jsii.String("targetRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopytarget.html
//
type CfnLifecyclePolicy_CrossRegionCopyTargetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopytarget.html#cfn-dlm-lifecyclepolicy-crossregioncopytarget-targetregion
	//
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}


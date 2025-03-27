package awspcs


// An Amazon EC2 launch template AWS PCS uses to launch compute nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLaunchTemplateProperty := &CustomLaunchTemplateProperty{
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	TemplateId: jsii.String("templateId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-customlaunchtemplate.html
//
type CfnComputeNodeGroup_CustomLaunchTemplateProperty struct {
	// The version of the EC2 launch template to use to provision instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-customlaunchtemplate.html#cfn-pcs-computenodegroup-customlaunchtemplate-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the EC2 launch template to use to provision instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-customlaunchtemplate.html#cfn-pcs-computenodegroup-customlaunchtemplate-templateid
	//
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}


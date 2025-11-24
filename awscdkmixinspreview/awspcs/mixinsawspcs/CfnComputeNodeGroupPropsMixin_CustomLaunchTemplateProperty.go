package mixinsawspcs


// An Amazon EC2 launch template AWS PCS uses to launch compute nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customLaunchTemplateProperty := &CustomLaunchTemplateProperty{
//   	TemplateId: jsii.String("templateId"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-customlaunchtemplate.html
//
type CfnComputeNodeGroupPropsMixin_CustomLaunchTemplateProperty struct {
	// The ID of the EC2 launch template to use to provision instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-customlaunchtemplate.html#cfn-pcs-computenodegroup-customlaunchtemplate-templateid
	//
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// The version of the EC2 launch template to use to provision instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-customlaunchtemplate.html#cfn-pcs-computenodegroup-customlaunchtemplate-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}


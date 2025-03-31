package awssystemsmanagersap


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentInfoProperty := &ComponentInfoProperty{
//   	ComponentType: jsii.String("componentType"),
//   	Ec2InstanceId: jsii.String("ec2InstanceId"),
//   	Sid: jsii.String("sid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html
//
type CfnApplication_ComponentInfoProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html#cfn-systemsmanagersap-application-componentinfo-componenttype
	//
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html#cfn-systemsmanagersap-application-componentinfo-ec2instanceid
	//
	Ec2InstanceId *string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html#cfn-systemsmanagersap-application-componentinfo-sid
	//
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}


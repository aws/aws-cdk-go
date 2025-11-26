package previewawssystemsmanagersapmixins


// This is information about the component of your SAP application, such as Web Dispatcher.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentInfoProperty := &ComponentInfoProperty{
//   	ComponentType: jsii.String("componentType"),
//   	Ec2InstanceId: jsii.String("ec2InstanceId"),
//   	Sid: jsii.String("sid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html
//
type CfnApplicationPropsMixin_ComponentInfoProperty struct {
	// This string is the type of the component.
	//
	// Accepted value is `WD` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html#cfn-systemsmanagersap-application-componentinfo-componenttype
	//
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// This is the Amazon EC2 instance on which your SAP component is running.
	//
	// Accepted values are alphanumeric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html#cfn-systemsmanagersap-application-componentinfo-ec2instanceid
	//
	Ec2InstanceId *string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// This string is the SAP System ID of the component.
	//
	// Accepted values are alphanumeric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-systemsmanagersap-application-componentinfo.html#cfn-systemsmanagersap-application-componentinfo-sid
	//
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}


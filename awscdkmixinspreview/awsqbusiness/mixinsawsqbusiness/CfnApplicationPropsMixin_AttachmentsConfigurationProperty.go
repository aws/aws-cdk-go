package mixinsawsqbusiness


// Configuration information for the file upload during chat feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attachmentsConfigurationProperty := &AttachmentsConfigurationProperty{
//   	AttachmentsControlMode: jsii.String("attachmentsControlMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-attachmentsconfiguration.html
//
type CfnApplicationPropsMixin_AttachmentsConfigurationProperty struct {
	// Status information about whether file upload functionality is activated or deactivated for your end user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-attachmentsconfiguration.html#cfn-qbusiness-application-attachmentsconfiguration-attachmentscontrolmode
	//
	AttachmentsControlMode *string `field:"optional" json:"attachmentsControlMode" yaml:"attachmentsControlMode"`
}


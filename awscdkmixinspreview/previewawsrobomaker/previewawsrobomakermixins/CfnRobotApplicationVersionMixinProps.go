package previewawsrobomakermixins


// Properties for CfnRobotApplicationVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRobotApplicationVersionMixinProps := &CfnRobotApplicationVersionMixinProps{
//   	Application: jsii.String("application"),
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html
//
type CfnRobotApplicationVersionMixinProps struct {
	// The application information for the robot application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html#cfn-robomaker-robotapplicationversion-application
	//
	Application *string `field:"optional" json:"application" yaml:"application"`
	// The current revision id for the robot application.
	//
	// If you provide a value and it matches the latest revision ID, a new version will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html#cfn-robomaker-robotapplicationversion-currentrevisionid
	//
	CurrentRevisionId *string `field:"optional" json:"currentRevisionId" yaml:"currentRevisionId"`
}


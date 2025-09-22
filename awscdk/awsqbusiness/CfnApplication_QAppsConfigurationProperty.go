package awsqbusiness


// Configuration information about Amazon Q Apps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qAppsConfigurationProperty := &QAppsConfigurationProperty{
//   	QAppsControlMode: jsii.String("qAppsControlMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-qappsconfiguration.html
//
type CfnApplication_QAppsConfigurationProperty struct {
	// Status information about whether end users can create and use Amazon Q Apps in the web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-qappsconfiguration.html#cfn-qbusiness-application-qappsconfiguration-qappscontrolmode
	//
	QAppsControlMode *string `field:"required" json:"qAppsControlMode" yaml:"qAppsControlMode"`
}


package awsconnect


// A first-party application's metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowModuleProperty := &FlowModuleProperty{
//   	FlowModuleId: jsii.String("flowModuleId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-flowmodule.html
//
type CfnSecurityProfile_FlowModuleProperty struct {
	// The identifier of the application that you want to give access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-flowmodule.html#cfn-connect-securityprofile-flowmodule-flowmoduleid
	//
	FlowModuleId *string `field:"required" json:"flowModuleId" yaml:"flowModuleId"`
	// The type of the first-party application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-flowmodule.html#cfn-connect-securityprofile-flowmodule-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}


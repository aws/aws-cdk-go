package previewawsconnectmixins


// This API is in preview release for Amazon Connect and is subject to change.
//
// A third-party application's metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationProperty := &ApplicationProperty{
//   	ApplicationPermissions: []*string{
//   		jsii.String("applicationPermissions"),
//   	},
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-application.html
//
type CfnSecurityProfilePropsMixin_ApplicationProperty struct {
	// The permissions that the agent is granted on the application.
	//
	// For third-party applications, only the `ACCESS` permission is supported. For MCP Servers, the permissions are tool Identifiers accepted by MCP Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-application.html#cfn-connect-securityprofile-application-applicationpermissions
	//
	ApplicationPermissions *[]*string `field:"optional" json:"applicationPermissions" yaml:"applicationPermissions"`
	// Namespace of the application that you want to give access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-application.html#cfn-connect-securityprofile-application-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}


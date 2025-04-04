package awsappsync


// Describes an authorization configuration.
//
// Use `AuthMode` to specify the publishing and subscription authorization configuration for an Event API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authModeProperty := &AuthModeProperty{
//   	AuthType: jsii.String("authType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-authmode.html
//
type CfnChannelNamespace_AuthModeProperty struct {
	// The authorization type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-authmode.html#cfn-appsync-channelnamespace-authmode-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
}


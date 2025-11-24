package mixinsawsappstream


// The streaming protocol that you want your stack to prefer.
//
// This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamingExperienceSettingsProperty := &StreamingExperienceSettingsProperty{
//   	PreferredProtocol: jsii.String("preferredProtocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-streamingexperiencesettings.html
//
type CfnStackPropsMixin_StreamingExperienceSettingsProperty struct {
	// The preferred protocol that you want to use while streaming your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-streamingexperiencesettings.html#cfn-appstream-stack-streamingexperiencesettings-preferredprotocol
	//
	PreferredProtocol *string `field:"optional" json:"preferredProtocol" yaml:"preferredProtocol"`
}


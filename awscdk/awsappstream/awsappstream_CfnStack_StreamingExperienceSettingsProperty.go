package awsappstream


// The streaming protocol that you want your stack to prefer.
//
// This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingExperienceSettingsProperty := &streamingExperienceSettingsProperty{
//   	preferredProtocol: jsii.String("preferredProtocol"),
//   }
//
type CfnStack_StreamingExperienceSettingsProperty struct {
	// The preferred protocol that you want to use while streaming your application.
	PreferredProtocol *string `field:"optional" json:"preferredProtocol" yaml:"preferredProtocol"`
}


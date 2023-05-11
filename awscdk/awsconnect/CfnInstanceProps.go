package awsconnect


// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &CfnInstanceProps{
//   	Attributes: &AttributesProperty{
//   		InboundCalls: jsii.Boolean(false),
//   		OutboundCalls: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AutoResolveBestVoices: jsii.Boolean(false),
//   		ContactflowLogs: jsii.Boolean(false),
//   		ContactLens: jsii.Boolean(false),
//   		EarlyMedia: jsii.Boolean(false),
//   		UseCustomTtsVoices: jsii.Boolean(false),
//   	},
//   	IdentityManagementType: jsii.String("identityManagementType"),
//
//   	// the properties below are optional
//   	DirectoryId: jsii.String("directoryId"),
//   	InstanceAlias: jsii.String("instanceAlias"),
//   }
//
type CfnInstanceProps struct {
	// A toggle for an individual feature at the instance level.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The identity management type.
	IdentityManagementType *string `field:"required" json:"identityManagementType" yaml:"identityManagementType"`
	// The identifier for the directory.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The alias of instance.
	//
	// `InstanceAlias` is only required when `IdentityManagementType` is `CONNECT_MANAGED` or `SAML` . `InstanceAlias` is not required when `IdentityManagementType` is `EXISTING_DIRECTORY` .
	InstanceAlias *string `field:"optional" json:"instanceAlias" yaml:"instanceAlias"`
}


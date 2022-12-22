package awsconnect


// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &cfnInstanceProps{
//   	attributes: &attributesProperty{
//   		inboundCalls: jsii.Boolean(false),
//   		outboundCalls: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		autoResolveBestVoices: jsii.Boolean(false),
//   		contactflowLogs: jsii.Boolean(false),
//   		contactLens: jsii.Boolean(false),
//   		earlyMedia: jsii.Boolean(false),
//   		useCustomTtsVoices: jsii.Boolean(false),
//   	},
//   	identityManagementType: jsii.String("identityManagementType"),
//
//   	// the properties below are optional
//   	directoryId: jsii.String("directoryId"),
//   	instanceAlias: jsii.String("instanceAlias"),
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
	InstanceAlias *string `field:"optional" json:"instanceAlias" yaml:"instanceAlias"`
}


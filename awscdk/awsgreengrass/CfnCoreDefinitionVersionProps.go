package awsgreengrass


// Properties for defining a `CfnCoreDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCoreDefinitionVersionProps := &CfnCoreDefinitionVersionProps{
//   	CoreDefinitionId: jsii.String("coreDefinitionId"),
//   	Cores: []interface{}{
//   		&CoreProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Id: jsii.String("id"),
//   			ThingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			SyncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html
//
type CfnCoreDefinitionVersionProps struct {
	// The ID of the core definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html#cfn-greengrass-coredefinitionversion-coredefinitionid
	//
	CoreDefinitionId *string `field:"required" json:"coreDefinitionId" yaml:"coreDefinitionId"`
	// The Greengrass core in this version.
	//
	// Currently, the `Cores` property for a core definition version can contain only one core.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html#cfn-greengrass-coredefinitionversion-cores
	//
	Cores interface{} `field:"required" json:"cores" yaml:"cores"`
}


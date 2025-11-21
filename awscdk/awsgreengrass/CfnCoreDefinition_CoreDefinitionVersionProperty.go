package awsgreengrass


// A core definition version contains a Greengrass [core](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html) .
//
// > After you create a core definition version that contains the core you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an CloudFormation template, `CoreDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::CoreDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreDefinitionVersionProperty := &CoreDefinitionVersionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-coredefinitionversion.html
//
type CfnCoreDefinition_CoreDefinitionVersionProperty struct {
	// The Greengrass core in this version.
	//
	// Currently, the `Cores` property for a core definition version can contain only one core.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-coredefinitionversion.html#cfn-greengrass-coredefinition-coredefinitionversion-cores
	//
	Cores interface{} `field:"required" json:"cores" yaml:"cores"`
}


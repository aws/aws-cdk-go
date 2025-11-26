package previewawsgreengrassmixins


// A group version in AWS IoT Greengrass , which references of a core definition version, device definition version, subscription definition version, and other version types that contain the components you want to deploy to a Greengrass core device.
//
// The group version must reference a core definition version that contains one core. Other version types are optionally included, depending on your business need.
//
// In an CloudFormation template, `GroupVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   groupVersionProperty := &GroupVersionProperty{
//   	ConnectorDefinitionVersionArn: jsii.String("connectorDefinitionVersionArn"),
//   	CoreDefinitionVersionArn: jsii.String("coreDefinitionVersionArn"),
//   	DeviceDefinitionVersionArn: jsii.String("deviceDefinitionVersionArn"),
//   	FunctionDefinitionVersionArn: jsii.String("functionDefinitionVersionArn"),
//   	LoggerDefinitionVersionArn: jsii.String("loggerDefinitionVersionArn"),
//   	ResourceDefinitionVersionArn: jsii.String("resourceDefinitionVersionArn"),
//   	SubscriptionDefinitionVersionArn: jsii.String("subscriptionDefinitionVersionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html
//
type CfnGroupPropsMixin_GroupVersionProperty struct {
	// The Amazon Resource Name (ARN) of the connector definition version that contains the connectors you want to deploy with the group version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-connectordefinitionversionarn
	//
	ConnectorDefinitionVersionArn *string `field:"optional" json:"connectorDefinitionVersionArn" yaml:"connectorDefinitionVersionArn"`
	// The ARN of the core definition version that contains the core you want to deploy with the group version.
	//
	// Currently, the core definition version can contain only one core.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-coredefinitionversionarn
	//
	CoreDefinitionVersionArn *string `field:"optional" json:"coreDefinitionVersionArn" yaml:"coreDefinitionVersionArn"`
	// The ARN of the device definition version that contains the devices you want to deploy with the group version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-devicedefinitionversionarn
	//
	DeviceDefinitionVersionArn *string `field:"optional" json:"deviceDefinitionVersionArn" yaml:"deviceDefinitionVersionArn"`
	// The ARN of the function definition version that contains the functions you want to deploy with the group version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-functiondefinitionversionarn
	//
	FunctionDefinitionVersionArn *string `field:"optional" json:"functionDefinitionVersionArn" yaml:"functionDefinitionVersionArn"`
	// The ARN of the logger definition version that contains the loggers you want to deploy with the group version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-loggerdefinitionversionarn
	//
	LoggerDefinitionVersionArn *string `field:"optional" json:"loggerDefinitionVersionArn" yaml:"loggerDefinitionVersionArn"`
	// The ARN of the resource definition version that contains the resources you want to deploy with the group version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-resourcedefinitionversionarn
	//
	ResourceDefinitionVersionArn *string `field:"optional" json:"resourceDefinitionVersionArn" yaml:"resourceDefinitionVersionArn"`
	// The ARN of the subscription definition version that contains the subscriptions you want to deploy with the group version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-group-groupversion.html#cfn-greengrass-group-groupversion-subscriptiondefinitionversionarn
	//
	SubscriptionDefinitionVersionArn *string `field:"optional" json:"subscriptionDefinitionVersionArn" yaml:"subscriptionDefinitionVersionArn"`
}


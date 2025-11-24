package mixinsawslightsail


// `Container` is a property of the [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html) property. It describes the settings of a container that will be launched, or that is launched, to an Amazon Lightsail container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerProperty := &ContainerProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	ContainerName: jsii.String("containerName"),
//   	Environment: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Value: jsii.String("value"),
//   			Variable: jsii.String("variable"),
//   		},
//   	},
//   	Image: jsii.String("image"),
//   	Ports: []interface{}{
//   		&PortInfoProperty{
//   			Port: jsii.String("port"),
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html
//
type CfnContainerPropsMixin_ContainerProperty struct {
	// The launch command for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html#cfn-lightsail-container-container-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html#cfn-lightsail-container-container-containername
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The environment variables of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html#cfn-lightsail-container-container-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The name of the image used for the container.
	//
	// Container images that are sourced from (registered and stored on) your container service start with a colon ( `:` ). For example, if your container service name is `container-service-1` , the container image label is `mystaticsite` , and you want to use the third version ( `3` ) of the registered container image, then you should specify `:container-service-1.mystaticsite.3` . To use the latest version of a container image, specify `latest` instead of a version number (for example, `:container-service-1.mystaticsite.latest` ). Your container service will automatically use the highest numbered version of the registered container image.
	//
	// Container images that are sourced from a public registry like Docker Hub don’t start with a colon. For example, `nginx:latest` or `nginx` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html#cfn-lightsail-container-container-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// An object that describes the open firewall ports and protocols of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html#cfn-lightsail-container-container-ports
	//
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}


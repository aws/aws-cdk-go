package awslightsail


// `Container` is a property of the [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html) property. It describes the settings of a container that will be launched, or that is launched, to an Amazon Lightsail container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerProperty := &containerProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	environment: []interface{}{
//   		&environmentVariableProperty{
//   			value: jsii.String("value"),
//   			variable: jsii.String("variable"),
//   		},
//   	},
//   	image: jsii.String("image"),
//   	ports: []interface{}{
//   		&portInfoProperty{
//   			port: jsii.String("port"),
//   			protocol: jsii.String("protocol"),
//   		},
//   	},
//   }
//
type CfnContainer_ContainerProperty struct {
	// The launch command for the container.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The environment variables of the container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The name of the image used for the container.
	//
	// Container images that are sourced from (registered and stored on) your container service start with a colon ( `:` ). For example, if your container service name is `container-service-1` , the container image label is `mystaticsite` , and you want to use the third version ( `3` ) of the registered container image, then you should specify `:container-service-1.mystaticsite.3` . To use the latest version of a container image, specify `latest` instead of a version number (for example, `:container-service-1.mystaticsite.latest` ). Your container service will automatically use the highest numbered version of the registered container image.
	//
	// Container images that are sourced from a public registry like Docker Hub don’t start with a colon. For example, `nginx:latest` or `nginx` .
	Image *string `field:"optional" json:"image" yaml:"image"`
	// An object that describes the open firewall ports and protocols of the container.
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}


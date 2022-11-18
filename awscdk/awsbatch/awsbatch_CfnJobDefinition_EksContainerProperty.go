package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var requests interface{}
//
//   eksContainerProperty := &eksContainerProperty{
//   	image: jsii.String("image"),
//
//   	// the properties below are optional
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	env: []interface{}{
//   		&eksContainerEnvironmentVariableProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   	imagePullPolicy: jsii.String("imagePullPolicy"),
//   	name: jsii.String("name"),
//   	resources: &resourcesProperty{
//   		limits: limits,
//   		requests: requests,
//   	},
//   	securityContext: &securityContextProperty{
//   		privileged: jsii.Boolean(false),
//   		readOnlyRootFilesystem: jsii.Boolean(false),
//   		runAsGroup: jsii.Number(123),
//   		runAsNonRoot: jsii.Boolean(false),
//   		runAsUser: jsii.Number(123),
//   	},
//   	volumeMounts: []interface{}{
//   		&eksContainerVolumeMountProperty{
//   			mountPath: jsii.String("mountPath"),
//   			name: jsii.String("name"),
//   			readOnly: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnJobDefinition_EksContainerProperty struct {
	// `CfnJobDefinition.EksContainerProperty.Image`.
	Image *string `field:"required" json:"image" yaml:"image"`
	// `CfnJobDefinition.EksContainerProperty.Args`.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// `CfnJobDefinition.EksContainerProperty.Command`.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// `CfnJobDefinition.EksContainerProperty.Env`.
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// `CfnJobDefinition.EksContainerProperty.ImagePullPolicy`.
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// `CfnJobDefinition.EksContainerProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnJobDefinition.EksContainerProperty.Resources`.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// `CfnJobDefinition.EksContainerProperty.SecurityContext`.
	SecurityContext interface{} `field:"optional" json:"securityContext" yaml:"securityContext"`
	// `CfnJobDefinition.EksContainerProperty.VolumeMounts`.
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}


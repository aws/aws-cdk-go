package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var requests interface{}
//
//   podPropertiesProperty := &podPropertiesProperty{
//   	containers: []interface{}{
//   		&eksContainerProperty{
//   			image: jsii.String("image"),
//
//   			// the properties below are optional
//   			args: []*string{
//   				jsii.String("args"),
//   			},
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			env: []interface{}{
//   				&eksContainerEnvironmentVariableProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   			imagePullPolicy: jsii.String("imagePullPolicy"),
//   			name: jsii.String("name"),
//   			resources: &resourcesProperty{
//   				limits: limits,
//   				requests: requests,
//   			},
//   			securityContext: &securityContextProperty{
//   				privileged: jsii.Boolean(false),
//   				readOnlyRootFilesystem: jsii.Boolean(false),
//   				runAsGroup: jsii.Number(123),
//   				runAsNonRoot: jsii.Boolean(false),
//   				runAsUser: jsii.Number(123),
//   			},
//   			volumeMounts: []interface{}{
//   				&eksContainerVolumeMountProperty{
//   					mountPath: jsii.String("mountPath"),
//   					name: jsii.String("name"),
//   					readOnly: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	dnsPolicy: jsii.String("dnsPolicy"),
//   	hostNetwork: jsii.Boolean(false),
//   	serviceAccountName: jsii.String("serviceAccountName"),
//   	volumes: []interface{}{
//   		&eksVolumeProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			emptyDir: &emptyDirProperty{
//   				medium: jsii.String("medium"),
//   				sizeLimit: jsii.String("sizeLimit"),
//   			},
//   			hostPath: &hostPathProperty{
//   				path: jsii.String("path"),
//   			},
//   			secret: &secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   }
//
type CfnJobDefinition_PodPropertiesProperty struct {
	// `CfnJobDefinition.PodPropertiesProperty.Containers`.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// `CfnJobDefinition.PodPropertiesProperty.DnsPolicy`.
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// `CfnJobDefinition.PodPropertiesProperty.HostNetwork`.
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// `CfnJobDefinition.PodPropertiesProperty.ServiceAccountName`.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// `CfnJobDefinition.PodPropertiesProperty.Volumes`.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}


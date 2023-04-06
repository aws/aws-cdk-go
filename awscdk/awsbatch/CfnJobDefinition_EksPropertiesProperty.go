package awsbatch


// An object that contains the properties for the Kubernetes resources of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var requests interface{}
//
//   eksPropertiesProperty := &EksPropertiesProperty{
//   	PodProperties: &PodPropertiesProperty{
//   		Containers: []interface{}{
//   			&EksContainerProperty{
//   				Image: jsii.String("image"),
//
//   				// the properties below are optional
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Env: []interface{}{
//   					&EksContainerEnvironmentVariableProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				ImagePullPolicy: jsii.String("imagePullPolicy"),
//   				Name: jsii.String("name"),
//   				Resources: &ResourcesProperty{
//   					Limits: limits,
//   					Requests: requests,
//   				},
//   				SecurityContext: &SecurityContextProperty{
//   					Privileged: jsii.Boolean(false),
//   					ReadOnlyRootFilesystem: jsii.Boolean(false),
//   					RunAsGroup: jsii.Number(123),
//   					RunAsNonRoot: jsii.Boolean(false),
//   					RunAsUser: jsii.Number(123),
//   				},
//   				VolumeMounts: []interface{}{
//   					&EksContainerVolumeMountProperty{
//   						MountPath: jsii.String("mountPath"),
//   						Name: jsii.String("name"),
//   						ReadOnly: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   		DnsPolicy: jsii.String("dnsPolicy"),
//   		HostNetwork: jsii.Boolean(false),
//   		ServiceAccountName: jsii.String("serviceAccountName"),
//   		Volumes: []interface{}{
//   			&EksVolumeProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				EmptyDir: &EmptyDirProperty{
//   					Medium: jsii.String("medium"),
//   					SizeLimit: jsii.String("sizeLimit"),
//   				},
//   				HostPath: &HostPathProperty{
//   					Path: jsii.String("path"),
//   				},
//   				Secret: &SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnJobDefinition_EksPropertiesProperty struct {
	// The properties for the Kubernetes pod resources of a job.
	PodProperties interface{} `field:"optional" json:"podProperties" yaml:"podProperties"`
}


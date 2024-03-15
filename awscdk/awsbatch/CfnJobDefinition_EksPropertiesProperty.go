package awsbatch


// An object that contains the properties for the Kubernetes resources of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var labels interface{}
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
//   		InitContainers: []interface{}{
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
//   		Metadata: &MetadataProperty{
//   			Labels: labels,
//   		},
//   		ServiceAccountName: jsii.String("serviceAccountName"),
//   		ShareProcessNamespace: jsii.Boolean(false),
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
//   				Secret: &EksSecretProperty{
//   					SecretName: jsii.String("secretName"),
//
//   					// the properties below are optional
//   					Optional: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksproperties.html
//
type CfnJobDefinition_EksPropertiesProperty struct {
	// The properties for the Kubernetes pod resources of a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksproperties.html#cfn-batch-jobdefinition-eksproperties-podproperties
	//
	PodProperties interface{} `field:"optional" json:"podProperties" yaml:"podProperties"`
}


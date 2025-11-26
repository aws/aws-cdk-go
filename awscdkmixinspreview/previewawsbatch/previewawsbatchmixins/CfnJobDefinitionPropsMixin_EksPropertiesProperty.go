package previewawsbatchmixins


// An object that contains the properties for the Kubernetes resources of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var labels interface{}
//   var limits interface{}
//   var requests interface{}
//
//   eksPropertiesProperty := &EksPropertiesProperty{
//   	PodProperties: &PodPropertiesProperty{
//   		Containers: []interface{}{
//   			&EksContainerProperty{
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Env: []interface{}{
//   					&EksContainerEnvironmentVariableProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Image: jsii.String("image"),
//   				ImagePullPolicy: jsii.String("imagePullPolicy"),
//   				Name: jsii.String("name"),
//   				Resources: &ResourcesProperty{
//   					Limits: limits,
//   					Requests: requests,
//   				},
//   				SecurityContext: &SecurityContextProperty{
//   					AllowPrivilegeEscalation: jsii.Boolean(false),
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
//   						SubPath: jsii.String("subPath"),
//   					},
//   				},
//   			},
//   		},
//   		DnsPolicy: jsii.String("dnsPolicy"),
//   		HostNetwork: jsii.Boolean(false),
//   		ImagePullSecrets: []interface{}{
//   			&ImagePullSecretProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		InitContainers: []interface{}{
//   			&EksContainerProperty{
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Env: []interface{}{
//   					&EksContainerEnvironmentVariableProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Image: jsii.String("image"),
//   				ImagePullPolicy: jsii.String("imagePullPolicy"),
//   				Name: jsii.String("name"),
//   				Resources: &ResourcesProperty{
//   					Limits: limits,
//   					Requests: requests,
//   				},
//   				SecurityContext: &SecurityContextProperty{
//   					AllowPrivilegeEscalation: jsii.Boolean(false),
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
//   						SubPath: jsii.String("subPath"),
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
//   				EmptyDir: &EmptyDirProperty{
//   					Medium: jsii.String("medium"),
//   					SizeLimit: jsii.String("sizeLimit"),
//   				},
//   				HostPath: &HostPathProperty{
//   					Path: jsii.String("path"),
//   				},
//   				Name: jsii.String("name"),
//   				PersistentVolumeClaim: &EksPersistentVolumeClaimProperty{
//   					ClaimName: jsii.String("claimName"),
//   					ReadOnly: jsii.Boolean(false),
//   				},
//   				Secret: &EksSecretProperty{
//   					Optional: jsii.Boolean(false),
//   					SecretName: jsii.String("secretName"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksproperties.html
//
type CfnJobDefinitionPropsMixin_EksPropertiesProperty struct {
	// The properties for the Kubernetes pod resources of a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksproperties.html#cfn-batch-jobdefinition-eksproperties-podproperties
	//
	PodProperties interface{} `field:"optional" json:"podProperties" yaml:"podProperties"`
}


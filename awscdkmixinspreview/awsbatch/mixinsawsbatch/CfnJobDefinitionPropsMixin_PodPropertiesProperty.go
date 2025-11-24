package mixinsawsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var labels interface{}
//   var limits interface{}
//   var requests interface{}
//
//   podPropertiesProperty := &PodPropertiesProperty{
//   	Containers: []interface{}{
//   		&EksContainerProperty{
//   			Args: []*string{
//   				jsii.String("args"),
//   			},
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Env: []interface{}{
//   				&EksContainerEnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Image: jsii.String("image"),
//   			ImagePullPolicy: jsii.String("imagePullPolicy"),
//   			Name: jsii.String("name"),
//   			Resources: &ResourcesProperty{
//   				Limits: limits,
//   				Requests: requests,
//   			},
//   			SecurityContext: &SecurityContextProperty{
//   				AllowPrivilegeEscalation: jsii.Boolean(false),
//   				Privileged: jsii.Boolean(false),
//   				ReadOnlyRootFilesystem: jsii.Boolean(false),
//   				RunAsGroup: jsii.Number(123),
//   				RunAsNonRoot: jsii.Boolean(false),
//   				RunAsUser: jsii.Number(123),
//   			},
//   			VolumeMounts: []interface{}{
//   				&EksContainerVolumeMountProperty{
//   					MountPath: jsii.String("mountPath"),
//   					Name: jsii.String("name"),
//   					ReadOnly: jsii.Boolean(false),
//   					SubPath: jsii.String("subPath"),
//   				},
//   			},
//   		},
//   	},
//   	DnsPolicy: jsii.String("dnsPolicy"),
//   	HostNetwork: jsii.Boolean(false),
//   	ImagePullSecrets: []interface{}{
//   		&ImagePullSecretProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	InitContainers: []interface{}{
//   		&EksContainerProperty{
//   			Args: []*string{
//   				jsii.String("args"),
//   			},
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Env: []interface{}{
//   				&EksContainerEnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Image: jsii.String("image"),
//   			ImagePullPolicy: jsii.String("imagePullPolicy"),
//   			Name: jsii.String("name"),
//   			Resources: &ResourcesProperty{
//   				Limits: limits,
//   				Requests: requests,
//   			},
//   			SecurityContext: &SecurityContextProperty{
//   				AllowPrivilegeEscalation: jsii.Boolean(false),
//   				Privileged: jsii.Boolean(false),
//   				ReadOnlyRootFilesystem: jsii.Boolean(false),
//   				RunAsGroup: jsii.Number(123),
//   				RunAsNonRoot: jsii.Boolean(false),
//   				RunAsUser: jsii.Number(123),
//   			},
//   			VolumeMounts: []interface{}{
//   				&EksContainerVolumeMountProperty{
//   					MountPath: jsii.String("mountPath"),
//   					Name: jsii.String("name"),
//   					ReadOnly: jsii.Boolean(false),
//   					SubPath: jsii.String("subPath"),
//   				},
//   			},
//   		},
//   	},
//   	Metadata: &MetadataProperty{
//   		Labels: labels,
//   	},
//   	ServiceAccountName: jsii.String("serviceAccountName"),
//   	ShareProcessNamespace: jsii.Boolean(false),
//   	Volumes: []interface{}{
//   		&EksVolumeProperty{
//   			EmptyDir: &EmptyDirProperty{
//   				Medium: jsii.String("medium"),
//   				SizeLimit: jsii.String("sizeLimit"),
//   			},
//   			HostPath: &HostPathProperty{
//   				Path: jsii.String("path"),
//   			},
//   			Name: jsii.String("name"),
//   			PersistentVolumeClaim: &EksPersistentVolumeClaimProperty{
//   				ClaimName: jsii.String("claimName"),
//   				ReadOnly: jsii.Boolean(false),
//   			},
//   			Secret: &EksSecretProperty{
//   				Optional: jsii.Boolean(false),
//   				SecretName: jsii.String("secretName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html
//
type CfnJobDefinitionPropsMixin_PodPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-containers
	//
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-dnspolicy
	//
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-hostnetwork
	//
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-imagepullsecrets
	//
	ImagePullSecrets interface{} `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-initcontainers
	//
	InitContainers interface{} `field:"optional" json:"initContainers" yaml:"initContainers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-serviceaccountname
	//
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-shareprocessnamespace
	//
	ShareProcessNamespace interface{} `field:"optional" json:"shareProcessNamespace" yaml:"shareProcessNamespace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-podproperties.html#cfn-batch-jobdefinition-podproperties-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}


package previewawsomicsmixins


// Use a container registry map to specify mappings between the ECR private repository and one or more upstream registries.
//
// For more information, see [Container images](https://docs.aws.amazon.com/omics/latest/dev/workflows-ecr.html) in the *AWS HealthOmics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerRegistryMapProperty := &ContainerRegistryMapProperty{
//   	ImageMappings: []interface{}{
//   		&ImageMappingProperty{
//   			DestinationImage: jsii.String("destinationImage"),
//   			SourceImage: jsii.String("sourceImage"),
//   		},
//   	},
//   	RegistryMappings: []interface{}{
//   		&RegistryMappingProperty{
//   			EcrAccountId: jsii.String("ecrAccountId"),
//   			EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   			UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   			UpstreamRepositoryPrefix: jsii.String("upstreamRepositoryPrefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-containerregistrymap.html
//
type CfnWorkflowVersionPropsMixin_ContainerRegistryMapProperty struct {
	// Image mappings specify path mappings between the ECR private repository and their corresponding external repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-containerregistrymap.html#cfn-omics-workflowversion-containerregistrymap-imagemappings
	//
	ImageMappings interface{} `field:"optional" json:"imageMappings" yaml:"imageMappings"`
	// Mapping that provides the ECR repository path where upstream container images are pulled and synchronized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-containerregistrymap.html#cfn-omics-workflowversion-containerregistrymap-registrymappings
	//
	RegistryMappings interface{} `field:"optional" json:"registryMappings" yaml:"registryMappings"`
}


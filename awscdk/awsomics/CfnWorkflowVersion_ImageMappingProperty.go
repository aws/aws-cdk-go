package awsomics


// Specifies image mappings that workflow tasks can use.
//
// For example, you can replace all the task references of a public image to use an equivalent image in your private ECR repository. You can use image mappings with upstream registries that don't support pull through cache. You need to manually synchronize the upstream registry with your private repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageMappingProperty := &ImageMappingProperty{
//   	DestinationImage: jsii.String("destinationImage"),
//   	SourceImage: jsii.String("sourceImage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-imagemapping.html
//
type CfnWorkflowVersion_ImageMappingProperty struct {
	// Specifies the URI of the corresponding image in the private ECR registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-imagemapping.html#cfn-omics-workflowversion-imagemapping-destinationimage
	//
	DestinationImage *string `field:"optional" json:"destinationImage" yaml:"destinationImage"`
	// Specifies the URI of the source image in the upstream registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-imagemapping.html#cfn-omics-workflowversion-imagemapping-sourceimage
	//
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
}


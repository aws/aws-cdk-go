package mixinsawsimagebuilder


// Settings that Image Builder uses to configure the ECR repository and the output container images that Amazon Inspector scans.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ecrConfigurationProperty := &EcrConfigurationProperty{
//   	ContainerTags: []*string{
//   		jsii.String("containerTags"),
//   	},
//   	RepositoryName: jsii.String("repositoryName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-ecrconfiguration.html
//
type CfnImagePropsMixin_EcrConfigurationProperty struct {
	// Tags for Image Builder to apply to the output container image that Amazon Inspector scans.
	//
	// Tags can help you identify and manage your scanned images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-ecrconfiguration.html#cfn-imagebuilder-image-ecrconfiguration-containertags
	//
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The name of the container repository that Amazon Inspector scans to identify findings for your container images.
	//
	// The name includes the path for the repository location. If you don’t provide this information, Image Builder creates a repository in your account named `image-builder-image-scanning-repository` for vulnerability scans of your output container images.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-ecrconfiguration.html#cfn-imagebuilder-image-ecrconfiguration-repositoryname
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}


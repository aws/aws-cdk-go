package cloudassemblyschema


// Where to publish docker images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageDestination := &dockerImageDestination{
//   	imageTag: jsii.String("imageTag"),
//   	repositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	region: jsii.String("region"),
//   }
//
type DockerImageDestination struct {
	// The role that needs to be assumed while publishing this asset.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
	// The region where this asset will need to be published.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tag of the image to publish.
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// Name of the ECR repository to publish to.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
}


package awsecr


// The details about the repository that are publicly visible in the Amazon ECR Public Gallery.
//
// For more information, see [Amazon ECR Public repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the *Amazon ECR Public User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryCatalogDataProperty := &RepositoryCatalogDataProperty{
//   	AboutText: jsii.String("aboutText"),
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	OperatingSystems: []*string{
//   		jsii.String("operatingSystems"),
//   	},
//   	RepositoryDescription: jsii.String("repositoryDescription"),
//   	UsageText: jsii.String("usageText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html
//
type CfnPublicRepository_RepositoryCatalogDataProperty struct {
	// The longform description of the contents of the repository.
	//
	// This text appears in the repository details on the Amazon ECR Public Gallery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-abouttext
	//
	AboutText *string `field:"optional" json:"aboutText" yaml:"aboutText"`
	// The architecture tags that are associated with the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-architectures
	//
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// The operating system tags that are associated with the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-operatingsystems
	//
	OperatingSystems *[]*string `field:"optional" json:"operatingSystems" yaml:"operatingSystems"`
	// The short description of the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-repositorydescription
	//
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// The longform usage details of the contents of the repository.
	//
	// The usage text provides context for users of the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-publicrepository-repositorycatalogdata.html#cfn-ecr-publicrepository-repositorycatalogdata-usagetext
	//
	UsageText *string `field:"optional" json:"usageText" yaml:"usageText"`
}


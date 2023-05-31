package awsecr


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
type CfnPublicRepository_RepositoryCatalogDataProperty struct {
	// `CfnPublicRepository.RepositoryCatalogDataProperty.AboutText`.
	AboutText *string `field:"optional" json:"aboutText" yaml:"aboutText"`
	// `CfnPublicRepository.RepositoryCatalogDataProperty.Architectures`.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// `CfnPublicRepository.RepositoryCatalogDataProperty.OperatingSystems`.
	OperatingSystems *[]*string `field:"optional" json:"operatingSystems" yaml:"operatingSystems"`
	// `CfnPublicRepository.RepositoryCatalogDataProperty.RepositoryDescription`.
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// `CfnPublicRepository.RepositoryCatalogDataProperty.UsageText`.
	UsageText *string `field:"optional" json:"usageText" yaml:"usageText"`
}


package awsecr


// Properties for looking up an existing Repository.
//
// Example:
//   // import using repository name
//   repositoryFromName := ecr.Repository_FromRepositoryName(this, jsii.String("ImportedRepoByName"), jsii.String("my-repo-name"))
//
//   // import using repository ARN
//   repositoryFromArn := ecr.Repository_FromRepositoryArn(this, jsii.String("ImportedRepoByArn"), jsii.String("arn:aws:ecr:us-east-1:123456789012:repository/my-repo-name"))
//
//   // import using repository lookup
//   // You have to provide either `repositoryArn` or `repositoryName` to lookup the repository
//   repositoryFromLookup := ecr.Repository_FromLookup(this, jsii.String("ImportedRepoByLookup"), &RepositoryLookupOptions{
//   	RepositoryArn: jsii.String("arn:aws:ecr:us-east-1:123456789012:repository/my-repo-name"),
//   	RepositoryName: jsii.String("my-repo-name"),
//   })
//
type RepositoryLookupOptions struct {
	// The ARN of the repository.
	// Default: - Do not filter on repository ARN.
	//
	RepositoryArn *string `field:"optional" json:"repositoryArn" yaml:"repositoryArn"`
	// The name of the repository.
	// Default: - Do not filter on repository name.
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}


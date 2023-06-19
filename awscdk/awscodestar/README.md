# AWS::CodeStar Construct Library

## GitHub Repository

To create a new GitHub Repository and commit the assets from S3 bucket into the repository after it is created:

```go
import codestar "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"


codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &GitHubRepositoryProps{
	Owner: jsii.String("aws"),
	RepositoryName: jsii.String("aws-cdk"),
	AccessToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token"), &SecretsManagerSecretOptions{
		JsonField: jsii.String("token"),
	}),
	ContentsBucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("bucket-name")),
	ContentsKey: jsii.String("import.zip"),
})
```

## Update or Delete the GitHubRepository

At this moment, updates to the `GitHubRepository` are not supported and the repository will not be deleted upon the deletion of the CloudFormation stack. You will need to update or delete the GitHub repository manually.

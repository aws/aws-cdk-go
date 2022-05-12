# AWS::CodeStar Construct Library

## GitHub Repository

To create a new GitHub Repository and commit the assets from S3 bucket into the repository after it is created:

```go
import codestar "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"


codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &gitHubRepositoryProps{
	owner: jsii.String("aws"),
	repositoryName: jsii.String("aws-cdk"),
	accessToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token"), &secretsManagerSecretOptions{
		jsonField: jsii.String("token"),
	}),
	contentsBucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("bucket-name")),
	contentsKey: jsii.String("import.zip"),
})
```

## Update or Delete the GitHubRepository

At this moment, updates to the `GitHubRepository` are not supported and the repository will not be deleted upon the deletion of the CloudFormation stack. You will need to update or delete the GitHub repository manually.

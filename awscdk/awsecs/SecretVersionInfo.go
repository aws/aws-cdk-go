package awsecs


// Specify the secret's version id or version stage.
//
// Example:
//   var secret secret
//   var dbSecret secret
//   var parameter stringParameter
//   var taskDefinition taskDefinition
//   var s3Bucket bucket
//
//
//   newContainer := taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(1024),
//   	Environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	EnvironmentFiles: []environmentFile{
//   		ecs.*environmentFile_FromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile_FromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	Secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret_fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret_fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret_fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.AddEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//   newContainer.AddSecret(jsii.String("API_KEY"), ecs.secret_FromSecretsManager(secret))
//   newContainer.AddSecret(jsii.String("DB_PASSWORD"), ecs.secret_FromSecretsManager(secret, jsii.String("password")))
//
type SecretVersionInfo struct {
	// version id of the secret.
	// Default: - use default version id.
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// version stage of the secret.
	// Default: - use default version stage.
	//
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}


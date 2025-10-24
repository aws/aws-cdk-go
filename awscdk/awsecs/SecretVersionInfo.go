package awsecs


// Specify the secret's version id or version stage.
//
// Example:
//   var secret Secret
//   var dbSecret Secret
//   var parameter StringParameter
//   var taskDefinition TaskDefinition
//   var s3Bucket Bucket
//
//
//   newContainer := taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(1024),
//   	Environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	EnvironmentFiles: []EnvironmentFile{
//   		ecs.EnvironmentFile_FromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.EnvironmentFile_FromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	Secrets: map[string]Secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.Secret_fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.Secret_fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.Secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.Secret_fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.AddEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//   newContainer.AddSecret(jsii.String("API_KEY"), ecs.Secret_FromSecretsManager(secret))
//   newContainer.AddSecret(jsii.String("DB_PASSWORD"), ecs.Secret_FromSecretsManager(secret, jsii.String("password")))
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


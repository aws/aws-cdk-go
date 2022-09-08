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
//   newContainer := taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   	environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	environmentFiles: []environmentFile{
//   		ecs.*environmentFile.fromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile.fromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret.fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret.fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.addEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//   newContainer.addSecret(jsii.String("API_KEY"), ecs.secret.fromSecretsManager(secret))
//   newContainer.addSecret(jsii.String("DB_PASSWORD"), ecs.secret.fromSecretsManager(secret, jsii.String("password")))
//
type SecretVersionInfo struct {
	// version id of the secret.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// version stage of the secret.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}


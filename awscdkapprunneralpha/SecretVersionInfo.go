// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Specify the secret's version id or version stage.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//
//   secret := secretsmanager.NewSecret(stack, jsii.String("Secret"))
//   parameter := ssm.StringParameter_FromSecureStringParameterAttributes(stack, jsii.String("Parameter"), &SecureStringParameterAttributes{
//   	ParameterName: jsii.String("/name"),
//   	Version: jsii.Number(1),
//   })
//
//   service := apprunner.NewService(stack, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   			EnvironmentSecrets: map[string]secret{
//   				"SECRET": apprunner.*secret_fromSecretsManager(secret),
//   				"PARAMETER": apprunner.*secret_fromSsmParameter(parameter),
//   				"SECRET_ID": apprunner.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   					"versionId": jsii.String("version-id"),
//   				}),
//   				"SECRET_STAGE": apprunner.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   					"versionStage": jsii.String("version-stage"),
//   				}),
//   			},
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
//   service.AddSecret(jsii.String("LATER_SECRET"), apprunner.secret_FromSecretsManager(secret, jsii.String("field")))
//
// Experimental.
type SecretVersionInfo struct {
	// version id of the secret.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// version stage of the secret.
	// Experimental.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}


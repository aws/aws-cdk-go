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
//   parameter := ssm.stringParameter.fromSecureStringParameterAttributes(stack, jsii.String("Parameter"), &secureStringParameterAttributes{
//   	parameterName: jsii.String("/name"),
//   	version: jsii.Number(1),
//   })
//
//   service := apprunner.NewService(stack, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   			environmentSecrets: map[string]secret{
//   				"SECRET": apprunner.*secret.fromSecretsManager(secret),
//   				"PARAMETER": apprunner.*secret.fromSsmParameter(parameter),
//   				"SECRET_ID": apprunner.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   					"versionId": jsii.String("version-id"),
//   				}),
//   				"SECRET_STAGE": apprunner.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   					"versionStage": jsii.String("version-stage"),
//   				}),
//   			},
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
//   service.addSecret(jsii.String("LATER_SECRET"), apprunner.secret.fromSecretsManager(secret, jsii.String("field")))
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


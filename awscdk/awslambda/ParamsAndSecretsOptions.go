package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Parameters and Secrets Extension configuration options.
//
// Example:
//   import sm "github.com/aws/aws-cdk-go/awscdk"
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//
//   secret := sm.NewSecret(this, jsii.String("Secret"))
//   parameter := ssm.NewStringParameter(this, jsii.String("Parameter"), &StringParameterProps{
//   	ParameterName: jsii.String("mySsmParameterName"),
//   	StringValue: jsii.String("mySsmParameterValue"),
//   })
//
//   paramsAndSecrets := lambda.ParamsAndSecretsLayerVersion_FromVersion(lambda.ParamsAndSecretsVersions_V1_0_103, &ParamsAndSecretsOptions{
//   	CacheSize: jsii.Number(500),
//   	LogLevel: lambda.ParamsAndSecretsLogLevel_DEBUG,
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Architecture: lambda.Architecture_ARM_64(),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	ParamsAndSecrets: ParamsAndSecrets,
//   })
//
//   secret.grantRead(lambdaFunction)
//   parameter.grantRead(lambdaFunction)
//
type ParamsAndSecretsOptions struct {
	// Whether the Parameters and Secrets Extension will cache parameters and secrets.
	// Default: true.
	//
	CacheEnabled *bool `field:"optional" json:"cacheEnabled" yaml:"cacheEnabled"`
	// The maximum number of secrets and parameters to cache.
	//
	// Must be a value
	// from 0 to 1000. A value of 0 means there is no caching.
	//
	// Note: This variable is ignored if parameterStoreTtl and secretsManagerTtl
	// are 0.
	// Default: 1000.
	//
	CacheSize *float64 `field:"optional" json:"cacheSize" yaml:"cacheSize"`
	// The port for the local HTTP server.
	//
	// Valid port numbers are 1 - 65535.
	// Default: 2773.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The level of logging provided by the Parameters and Secrets Extension.
	//
	// Note: Set to debug to see the cache configuration.
	// Default: - Logging level will be `info`.
	//
	LogLevel ParamsAndSecretsLogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The maximum number of connection for HTTP clients that the Parameters and Secrets Extension uses to make requests to Parameter Store or Secrets Manager.
	//
	// There is no maximum limit. Minimum is 1.
	//
	// Note: Every running copy of this Lambda function may open the number of
	// connections specified by this property. Thus, the total number of connections
	// may exceed this number.
	// Default: 3.
	//
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// The timeout for requests to Parameter Store.
	//
	// A value of 0 means that there is no
	// timeout.
	// Default: 0.
	//
	ParameterStoreTimeout awscdk.Duration `field:"optional" json:"parameterStoreTimeout" yaml:"parameterStoreTimeout"`
	// The time-to-live of a parameter in the cache.
	//
	// A value of 0 means there is no caching.
	// The maximum time-to-live is 300 seconds.
	//
	// Note: This variable is ignored if cacheSize is 0.
	// Default: 300 seconds.
	//
	ParameterStoreTtl awscdk.Duration `field:"optional" json:"parameterStoreTtl" yaml:"parameterStoreTtl"`
	// The timeout for requests to Secrets Manager.
	//
	// A value of 0 means that there is
	// no timeout.
	// Default: 0.
	//
	SecretsManagerTimeout awscdk.Duration `field:"optional" json:"secretsManagerTimeout" yaml:"secretsManagerTimeout"`
	// The time-to-live of a secret in the cache.
	//
	// A value of 0 means there is no caching.
	// The maximum time-to-live is 300 seconds.
	//
	// Note: This variable is ignored if cacheSize is 0.
	// Default: 300 seconds.
	//
	SecretsManagerTtl awscdk.Duration `field:"optional" json:"secretsManagerTtl" yaml:"secretsManagerTtl"`
}


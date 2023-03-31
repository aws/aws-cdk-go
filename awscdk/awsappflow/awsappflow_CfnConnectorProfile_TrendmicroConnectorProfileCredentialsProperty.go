package awsappflow


// The connector-specific profile credentials required when using Trend Micro.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trendmicroConnectorProfileCredentialsProperty := &trendmicroConnectorProfileCredentialsProperty{
//   	apiSecretKey: jsii.String("apiSecretKey"),
//   }
//
type CfnConnectorProfile_TrendmicroConnectorProfileCredentialsProperty struct {
	// The Secret Access Key portion of the credentials.
	ApiSecretKey *string `field:"required" json:"apiSecretKey" yaml:"apiSecretKey"`
}


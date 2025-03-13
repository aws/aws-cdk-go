package awsappflow


// The connector-specific profile credentials required when using Trend Micro.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trendmicroConnectorProfileCredentialsProperty := &TrendmicroConnectorProfileCredentialsProperty{
//   	ApiSecretKey: jsii.String("apiSecretKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-trendmicroconnectorprofilecredentials.html
//
type CfnConnectorProfile_TrendmicroConnectorProfileCredentialsProperty struct {
	// The Secret Access Key portion of the credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-trendmicroconnectorprofilecredentials.html#cfn-appflow-connectorprofile-trendmicroconnectorprofilecredentials-apisecretkey
	//
	ApiSecretKey *string `field:"required" json:"apiSecretKey" yaml:"apiSecretKey"`
}


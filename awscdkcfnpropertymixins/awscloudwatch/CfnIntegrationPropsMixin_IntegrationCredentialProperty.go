package awscloudwatch


// Write-only API-key credential material for the integration.
//
// Stored encrypted and never returned on read.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   integrationCredentialProperty := &IntegrationCredentialProperty{
//   	ApiKeyCredential: &ApiKeyCredentialProperty{
//   		ApiKeyValue: jsii.String("apiKeyValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-integration-integrationcredential.html
//
type CfnIntegrationPropsMixin_IntegrationCredentialProperty struct {
	// API-key credential material that can be rotated in place on update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-integration-integrationcredential.html#cfn-cloudwatch-integration-integrationcredential-apikeycredential
	//
	ApiKeyCredential interface{} `field:"optional" json:"apiKeyCredential" yaml:"apiKeyCredential"`
}


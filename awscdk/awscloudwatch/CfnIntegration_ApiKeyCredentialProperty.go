package awscloudwatch


// API-key credential material that can be rotated in place on update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyCredentialProperty := &ApiKeyCredentialProperty{
//   	ApiKeyValue: jsii.String("apiKeyValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-integration-apikeycredential.html
//
type CfnIntegration_ApiKeyCredentialProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-integration-apikeycredential.html#cfn-cloudwatch-integration-apikeycredential-apikeyvalue
	//
	ApiKeyValue *string `field:"required" json:"apiKeyValue" yaml:"apiKeyValue"`
}


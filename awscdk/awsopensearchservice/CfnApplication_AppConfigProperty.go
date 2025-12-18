package awsopensearchservice


// Configuration settings for an OpenSearch application.
//
// For more information, see [Using the OpenSearch user interface in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/application.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appConfigProperty := &AppConfigProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-appconfig.html
//
type CfnApplication_AppConfigProperty struct {
	// The configuration item to set, such as the admin role for the OpenSearch application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-appconfig.html#cfn-opensearchservice-application-appconfig-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value assigned to the configuration key, such as an IAM user ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-appconfig.html#cfn-opensearchservice-application-appconfig-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


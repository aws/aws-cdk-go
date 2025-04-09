package awsopensearchservice


// Configurations of the OpenSearch Application.
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
	// Specify the item to configure, such as admin role for the OpenSearch Application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-appconfig.html#cfn-opensearchservice-application-appconfig-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specifies the value to configure for the key, such as an IAM user ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-appconfig.html#cfn-opensearchservice-application-appconfig-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


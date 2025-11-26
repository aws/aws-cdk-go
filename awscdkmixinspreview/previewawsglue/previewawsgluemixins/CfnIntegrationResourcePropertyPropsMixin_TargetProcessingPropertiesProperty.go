package previewawsgluemixins


// The structure used to define the resource properties associated with the integration target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetProcessingPropertiesProperty := &TargetProcessingPropertiesProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	EventBusArn: jsii.String("eventBusArn"),
//   	KmsArn: jsii.String("kmsArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-targetprocessingproperties.html
//
type CfnIntegrationResourcePropertyPropsMixin_TargetProcessingPropertiesProperty struct {
	// The AWS Glue network connection to configure the AWS Glue job running in the customer VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-targetprocessingproperties.html#cfn-glue-integrationresourceproperty-targetprocessingproperties-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The ARN of an Eventbridge event bus to receive the integration status notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-targetprocessingproperties.html#cfn-glue-integrationresourceproperty-targetprocessingproperties-eventbusarn
	//
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
	// The ARN of the KMS key used for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-targetprocessingproperties.html#cfn-glue-integrationresourceproperty-targetprocessingproperties-kmsarn
	//
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// The IAM role to access the AWS Glue database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-targetprocessingproperties.html#cfn-glue-integrationresourceproperty-targetprocessingproperties-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


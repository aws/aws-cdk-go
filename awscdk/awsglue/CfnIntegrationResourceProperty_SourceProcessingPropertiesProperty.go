package awsglue


// The structure used to define the resource properties associated with the integration source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProcessingPropertiesProperty := &SourceProcessingPropertiesProperty{
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-sourceprocessingproperties.html
//
type CfnIntegrationResourceProperty_SourceProcessingPropertiesProperty struct {
	// The IAM role to access the AWS Glue connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integrationresourceproperty-sourceprocessingproperties.html#cfn-glue-integrationresourceproperty-sourceprocessingproperties-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}


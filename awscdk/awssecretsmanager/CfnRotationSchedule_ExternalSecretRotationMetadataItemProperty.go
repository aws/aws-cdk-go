package awssecretsmanager


// The metadata needed to successfully rotate a managed external secret.
//
// Each metadata item is a key and value pair of strings in a JSON text string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   externalSecretRotationMetadataItemProperty := &ExternalSecretRotationMetadataItemProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.html
//
type CfnRotationSchedule_ExternalSecretRotationMetadataItemProperty struct {
	// The key name of the metadata item.
	//
	// You can specify a value that's 1 to 256 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.html#cfn-secretsmanager-rotationschedule-externalsecretrotationmetadataitem-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the metadata item.
	//
	// You can specify a value that's 1 to 2048 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.html#cfn-secretsmanager-rotationschedule-externalsecretrotationmetadataitem-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


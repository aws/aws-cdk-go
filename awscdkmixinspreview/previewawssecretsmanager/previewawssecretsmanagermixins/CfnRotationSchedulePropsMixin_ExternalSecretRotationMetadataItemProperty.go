package previewawssecretsmanagermixins


// The metadata needed to successfully rotate a managed external secret.
//
// A list of key value pairs in JSON format specified by the partner. For more information, see [Managed external secret partners](https://docs.aws.amazon.com/secretsmanager/latest/userguide/mes-partners.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   externalSecretRotationMetadataItemProperty := &ExternalSecretRotationMetadataItemProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.html
//
type CfnRotationSchedulePropsMixin_ExternalSecretRotationMetadataItemProperty struct {
	// The key that identifies the item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.html#cfn-secretsmanager-rotationschedule-externalsecretrotationmetadataitem-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the specified item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.html#cfn-secretsmanager-rotationschedule-externalsecretrotationmetadataitem-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


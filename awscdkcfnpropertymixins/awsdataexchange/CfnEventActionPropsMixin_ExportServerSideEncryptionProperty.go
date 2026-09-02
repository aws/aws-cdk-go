package awsdataexchange


// Encryption configuration of the export job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   exportServerSideEncryptionProperty := &ExportServerSideEncryptionProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-exportserversideencryption.html
//
type CfnEventActionPropsMixin_ExportServerSideEncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the AWS KMS key you want to use to encrypt the Amazon S3 objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-exportserversideencryption.html#cfn-dataexchange-eventaction-exportserversideencryption-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The type of server side encryption used for encrypting the objects in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-exportserversideencryption.html#cfn-dataexchange-eventaction-exportserversideencryption-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


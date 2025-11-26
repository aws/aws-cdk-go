package previewawsssmincidentsmixins


// The `RegionConfiguration` property specifies the Region and AWS Key Management Service key to add to the replication set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   regionConfigurationProperty := &RegionConfigurationProperty{
//   	SseKmsKeyId: jsii.String("sseKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-replicationset-regionconfiguration.html
//
type CfnReplicationSetPropsMixin_RegionConfigurationProperty struct {
	// The AWS Key Management Service key ID to use to encrypt your replication set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-replicationset-regionconfiguration.html#cfn-ssmincidents-replicationset-regionconfiguration-ssekmskeyid
	//
	SseKmsKeyId *string `field:"optional" json:"sseKmsKeyId" yaml:"sseKmsKeyId"`
}


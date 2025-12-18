package previewawspaymentcryptographymixins


// Represents the replication status information for a key in a replication region for [Multi-Region key replication](https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html) .
//
// This structure contains details about the current state of key replication, including any status messages and operational information about the replication process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationStatusTypeProperty := &ReplicationStatusTypeProperty{
//   	Status: jsii.String("status"),
//   	StatusMessage: jsii.String("statusMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-replicationstatustype.html
//
type CfnKeyPropsMixin_ReplicationStatusTypeProperty struct {
	// The current status of key replication in this AWS Region .
	//
	// This field indicates whether the key replication is in progress, completed successfully, or has encountered an error. Possible values include states such as `SYNCRHONIZED` , `IN_PROGRESS` , `DELETE_IN_PROGRESS` , or `FAILED` . This provides visibility into the replication process for monitoring and troubleshooting purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-replicationstatustype.html#cfn-paymentcryptography-key-replicationstatustype-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A message that provides additional information about the current replication status of the key.
	//
	// This field contains details about any issues or progress updates related to key replication operations. It may include information about replication failures, synchronization status, or other operational details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-paymentcryptography-key-replicationstatustype.html#cfn-paymentcryptography-key-replicationstatustype-statusmessage
	//
	StatusMessage *string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}


package awsdsql


// Configuration details about encryption for the cluster including the AWS  key ARN, encryption type, and encryption status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionDetailsProperty := &EncryptionDetailsProperty{
//   	EncryptionStatus: jsii.String("encryptionStatus"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-encryptiondetails.html
//
type CfnCluster_EncryptionDetailsProperty struct {
	// The status of encryption for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-encryptiondetails.html#cfn-dsql-cluster-encryptiondetails-encryptionstatus
	//
	EncryptionStatus *string `field:"optional" json:"encryptionStatus" yaml:"encryptionStatus"`
	// The type of encryption that protects the data on your cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-encryptiondetails.html#cfn-dsql-cluster-encryptiondetails-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The ARN of the AWS  key that encrypts data in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-encryptiondetails.html#cfn-dsql-cluster-encryptiondetails-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


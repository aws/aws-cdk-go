package awssecuritylake


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitionsProperty := &TransitionsProperty{
//   	Days: jsii.Number(123),
//   	StorageClass: jsii.String("storageClass"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-transitions.html
//
type CfnDataLake_TransitionsProperty struct {
	// Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-transitions.html#cfn-securitylake-datalake-transitions-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-transitions.html#cfn-securitylake-datalake-transitions-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}


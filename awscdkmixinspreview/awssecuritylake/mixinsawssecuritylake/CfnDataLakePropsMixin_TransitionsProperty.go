package mixinsawssecuritylake


// Provides transition lifecycle details of the Amazon Security Lake object.
//
// For more information about Amazon S3 Lifecycle configurations, see [Managing your storage lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html) in the *Amazon Simple Storage Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transitionsProperty := &TransitionsProperty{
//   	Days: jsii.Number(123),
//   	StorageClass: jsii.String("storageClass"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-transitions.html
//
type CfnDataLakePropsMixin_TransitionsProperty struct {
	// The number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-transitions.html#cfn-securitylake-datalake-transitions-days
	//
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The list of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
	//
	// The default storage class is *S3 Standard* . For information about other storage classes, see [Setting the storage class of an object](https://docs.aws.amazon.com/AmazonS3/latest/userguide/sc-howtoset.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-transitions.html#cfn-securitylake-datalake-transitions-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}


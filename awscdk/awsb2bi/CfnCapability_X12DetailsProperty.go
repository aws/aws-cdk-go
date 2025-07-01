package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12DetailsProperty := &X12DetailsProperty{
//   	TransactionSet: jsii.String("transactionSet"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-x12details.html
//
type CfnCapability_X12DetailsProperty struct {
	// Returns an enumerated type where each value identifies an X12 transaction set.
	//
	// Transaction sets are maintained by the X12 Accredited Standards Committee.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-x12details.html#cfn-b2bi-capability-x12details-transactionset
	//
	TransactionSet *string `field:"optional" json:"transactionSet" yaml:"transactionSet"`
	// Returns the version to use for the specified X12 transaction set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-x12details.html#cfn-b2bi-capability-x12details-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}


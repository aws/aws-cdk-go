package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ediTypeProperty := &EdiTypeProperty{
//   	X12Details: &X12DetailsProperty{
//   		TransactionSet: jsii.String("transactionSet"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-editype.html
//
type CfnCapability_EdiTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-capability-editype.html#cfn-b2bi-capability-editype-x12details
	//
	X12Details interface{} `field:"required" json:"x12Details" yaml:"x12Details"`
}


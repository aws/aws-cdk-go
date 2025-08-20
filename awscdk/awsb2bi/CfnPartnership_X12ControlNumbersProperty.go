package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12ControlNumbersProperty := &X12ControlNumbersProperty{
//   	StartingFunctionalGroupControlNumber: jsii.Number(123),
//   	StartingInterchangeControlNumber: jsii.Number(123),
//   	StartingTransactionSetControlNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html
//
type CfnPartnership_X12ControlNumbersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html#cfn-b2bi-partnership-x12controlnumbers-startingfunctionalgroupcontrolnumber
	//
	StartingFunctionalGroupControlNumber *float64 `field:"optional" json:"startingFunctionalGroupControlNumber" yaml:"startingFunctionalGroupControlNumber"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html#cfn-b2bi-partnership-x12controlnumbers-startinginterchangecontrolnumber
	//
	StartingInterchangeControlNumber *float64 `field:"optional" json:"startingInterchangeControlNumber" yaml:"startingInterchangeControlNumber"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html#cfn-b2bi-partnership-x12controlnumbers-startingtransactionsetcontrolnumber
	//
	StartingTransactionSetControlNumber *float64 `field:"optional" json:"startingTransactionSetControlNumber" yaml:"startingTransactionSetControlNumber"`
}


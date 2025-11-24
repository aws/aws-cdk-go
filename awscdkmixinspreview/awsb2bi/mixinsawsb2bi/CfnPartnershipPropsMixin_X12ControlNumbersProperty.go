package mixinsawsb2bi


// Contains configuration for X12 control numbers used in X12 EDI generation.
//
// Control numbers are used to uniquely identify interchanges, functional groups, and transaction sets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12ControlNumbersProperty := &X12ControlNumbersProperty{
//   	StartingFunctionalGroupControlNumber: jsii.Number(123),
//   	StartingInterchangeControlNumber: jsii.Number(123),
//   	StartingTransactionSetControlNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html
//
type CfnPartnershipPropsMixin_X12ControlNumbersProperty struct {
	// Specifies the starting functional group control number (GS06) to use for X12 EDI generation.
	//
	// This number is incremented for each new functional group. For the GS (functional group) envelope, AWS B2B Data Interchange generates a functional group control number that is unique to the sender ID, receiver ID, and functional identifier code combination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html#cfn-b2bi-partnership-x12controlnumbers-startingfunctionalgroupcontrolnumber
	//
	StartingFunctionalGroupControlNumber *float64 `field:"optional" json:"startingFunctionalGroupControlNumber" yaml:"startingFunctionalGroupControlNumber"`
	// Specifies the starting interchange control number (ISA13) to use for X12 EDI generation.
	//
	// This number is incremented for each new interchange. For the ISA (interchange) envelope, AWS B2B Data Interchange generates an interchange control number that is unique for the ISA05 and ISA06 (sender) & ISA07 and ISA08 (receiver) combination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html#cfn-b2bi-partnership-x12controlnumbers-startinginterchangecontrolnumber
	//
	StartingInterchangeControlNumber *float64 `field:"optional" json:"startingInterchangeControlNumber" yaml:"startingInterchangeControlNumber"`
	// Specifies the starting transaction set control number (ST02) to use for X12 EDI generation.
	//
	// This number is incremented for each new transaction set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12controlnumbers.html#cfn-b2bi-partnership-x12controlnumbers-startingtransactionsetcontrolnumber
	//
	StartingTransactionSetControlNumber *float64 `field:"optional" json:"startingTransactionSetControlNumber" yaml:"startingTransactionSetControlNumber"`
}


package awsconnect


// Identifier for hours of operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationsIdentifierProperty := &HoursOfOperationsIdentifierProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationsidentifier.html
//
type CfnHoursOfOperation_HoursOfOperationsIdentifierProperty struct {
	// The identifier for the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationsidentifier.html#cfn-connect-hoursofoperation-hoursofoperationsidentifier-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The name of the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationsidentifier.html#cfn-connect-hoursofoperation-hoursofoperationsidentifier-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


package awsappflow


// SAP Source connector page size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataPaginationConfigProperty := &SAPODataPaginationConfigProperty{
//   	MaxPageSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatapaginationconfig.html
//
type CfnFlow_SAPODataPaginationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatapaginationconfig.html#cfn-appflow-flow-sapodatapaginationconfig-maxpagesize
	//
	MaxPageSize *float64 `field:"required" json:"maxPageSize" yaml:"maxPageSize"`
}


package mixinsawsappflow


// Sets the page size for each *concurrent process* that transfers OData records from your SAP instance.
//
// A concurrent process is query that retrieves a batch of records as part of a flow run. Amazon AppFlow can run multiple concurrent processes in parallel to transfer data faster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sAPODataPaginationConfigProperty := &SAPODataPaginationConfigProperty{
//   	MaxPageSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatapaginationconfig.html
//
type CfnFlowPropsMixin_SAPODataPaginationConfigProperty struct {
	// The maximum number of records that Amazon AppFlow receives in each page of the response from your SAP application.
	//
	// For transfers of OData records, the maximum page size is 3,000. For transfers of data that comes from an ODP provider, the maximum page size is 10,000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodatapaginationconfig.html#cfn-appflow-flow-sapodatapaginationconfig-maxpagesize
	//
	MaxPageSize *float64 `field:"optional" json:"maxPageSize" yaml:"maxPageSize"`
}


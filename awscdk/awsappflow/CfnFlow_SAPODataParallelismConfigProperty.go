package awsappflow


// Sets the number of *concurrent processes* that transfer OData records from your SAP instance.
//
// A concurrent process is query that retrieves a batch of records as part of a flow run. Amazon AppFlow can run multiple concurrent processes in parallel to transfer data faster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataParallelismConfigProperty := &SAPODataParallelismConfigProperty{
//   	MaxParallelism: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodataparallelismconfig.html
//
type CfnFlow_SAPODataParallelismConfigProperty struct {
	// The maximum number of processes that Amazon AppFlow runs at the same time when it retrieves your data from your SAP application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodataparallelismconfig.html#cfn-appflow-flow-sapodataparallelismconfig-maxparallelism
	//
	MaxParallelism *float64 `field:"required" json:"maxParallelism" yaml:"maxParallelism"`
}


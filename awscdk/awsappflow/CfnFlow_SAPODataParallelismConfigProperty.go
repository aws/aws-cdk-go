package awsappflow


// SAP Source connector parallelism factor.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sapodataparallelismconfig.html#cfn-appflow-flow-sapodataparallelismconfig-maxparallelism
	//
	MaxParallelism *float64 `field:"required" json:"maxParallelism" yaml:"maxParallelism"`
}


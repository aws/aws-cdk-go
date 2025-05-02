package awsxray


// Properties for defining a `CfnTransactionSearchConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransactionSearchConfigProps := &CfnTransactionSearchConfigProps{
//   	IndexingPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-transactionsearchconfig.html
//
type CfnTransactionSearchConfigProps struct {
	// Determines the percentage of traces indexed from CloudWatch Logs to X-Ray.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-transactionsearchconfig.html#cfn-xray-transactionsearchconfig-indexingpercentage
	//
	IndexingPercentage *float64 `field:"optional" json:"indexingPercentage" yaml:"indexingPercentage"`
}


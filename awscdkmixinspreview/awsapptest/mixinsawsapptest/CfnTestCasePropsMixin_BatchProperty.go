package mixinsawsapptest


// Defines a batch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   batchProperty := &BatchProperty{
//   	BatchJobName: jsii.String("batchJobName"),
//   	BatchJobParameters: map[string]*string{
//   		"batchJobParametersKey": jsii.String("batchJobParameters"),
//   	},
//   	ExportDataSetNames: []*string{
//   		jsii.String("exportDataSetNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-batch.html
//
type CfnTestCasePropsMixin_BatchProperty struct {
	// The job name of the batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-batch.html#cfn-apptest-testcase-batch-batchjobname
	//
	BatchJobName *string `field:"optional" json:"batchJobName" yaml:"batchJobName"`
	// The batch job parameters of the batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-batch.html#cfn-apptest-testcase-batch-batchjobparameters
	//
	BatchJobParameters interface{} `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// The export data set names of the batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-batch.html#cfn-apptest-testcase-batch-exportdatasetnames
	//
	ExportDataSetNames *[]*string `field:"optional" json:"exportDataSetNames" yaml:"exportDataSetNames"`
}


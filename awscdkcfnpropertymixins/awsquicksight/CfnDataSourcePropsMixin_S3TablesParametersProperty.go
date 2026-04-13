package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3TablesParametersProperty := &S3TablesParametersProperty{
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-s3tablesparameters.html
//
type CfnDataSourcePropsMixin_S3TablesParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-s3tablesparameters.html#cfn-quicksight-datasource-s3tablesparameters-tablebucketarn
	//
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
}


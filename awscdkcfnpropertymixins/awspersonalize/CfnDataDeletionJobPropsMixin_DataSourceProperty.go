package awspersonalize


// The Amazon S3 bucket that contains the list of userIds to delete.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataSourceProperty := &DataSourceProperty{
//   	DataLocation: jsii.String("dataLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-datadeletionjob-datasource.html
//
type CfnDataDeletionJobPropsMixin_DataSourceProperty struct {
	// The path to the Amazon S3 bucket where the data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-datadeletionjob-datasource.html#cfn-personalize-datadeletionjob-datasource-datalocation
	//
	DataLocation *string `field:"optional" json:"dataLocation" yaml:"dataLocation"`
}


package awspersonalize


// Properties for CfnDataDeletionJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDataDeletionJobMixinProps := &CfnDataDeletionJobMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	DataSource: &DataSourceProperty{
//   		DataLocation: jsii.String("dataLocation"),
//   	},
//   	JobName: jsii.String("jobName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html
//
type CfnDataDeletionJobMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-jobname
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-datadeletionjob.html#cfn-personalize-datadeletionjob-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


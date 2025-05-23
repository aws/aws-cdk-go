package awsdatazone


// The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftStorageProperty := &RedshiftStorageProperty{
//   	RedshiftClusterSource: &RedshiftClusterStorageProperty{
//   		ClusterName: jsii.String("clusterName"),
//   	},
//   	RedshiftServerlessSource: &RedshiftServerlessStorageProperty{
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftstorage.html
//
type CfnDataSource_RedshiftStorageProperty struct {
	// The details of the Amazon Redshift cluster source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftstorage.html#cfn-datazone-datasource-redshiftstorage-redshiftclustersource
	//
	RedshiftClusterSource interface{} `field:"optional" json:"redshiftClusterSource" yaml:"redshiftClusterSource"`
	// The details of the Amazon Redshift Serverless workgroup source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftstorage.html#cfn-datazone-datasource-redshiftstorage-redshiftserverlesssource
	//
	RedshiftServerlessSource interface{} `field:"optional" json:"redshiftServerlessSource" yaml:"redshiftServerlessSource"`
}


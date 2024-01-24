package awsdatazone


// The name of an Amazon Redshift cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftClusterStorageProperty := &RedshiftClusterStorageProperty{
//   	ClusterName: jsii.String("clusterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftclusterstorage.html
//
type CfnDataSource_RedshiftClusterStorageProperty struct {
	// The name of an Amazon Redshift cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftclusterstorage.html#cfn-datazone-datasource-redshiftclusterstorage-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}


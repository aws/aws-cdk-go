package mixinsawsdatazone


// The details of the Amazon Redshift cluster storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftClusterStorageProperty := &RedshiftClusterStorageProperty{
//   	ClusterName: jsii.String("clusterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftclusterstorage.html
//
type CfnDataSourcePropsMixin_RedshiftClusterStorageProperty struct {
	// The name of an Amazon Redshift cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftclusterstorage.html#cfn-datazone-datasource-redshiftclusterstorage-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
}


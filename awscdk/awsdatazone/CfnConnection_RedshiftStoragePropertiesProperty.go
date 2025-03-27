package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftStoragePropertiesProperty := &RedshiftStoragePropertiesProperty{
//   	ClusterName: jsii.String("clusterName"),
//   	WorkgroupName: jsii.String("workgroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftstorageproperties.html
//
type CfnConnection_RedshiftStoragePropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftstorageproperties.html#cfn-datazone-connection-redshiftstorageproperties-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftstorageproperties.html#cfn-datazone-connection-redshiftstorageproperties-workgroupname
	//
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}


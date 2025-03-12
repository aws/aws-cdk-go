package awsdatazone


// The details of the Amazon Redshift Serverless workgroup storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftServerlessStorageProperty := &RedshiftServerlessStorageProperty{
//   	WorkgroupName: jsii.String("workgroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftserverlessstorage.html
//
type CfnDataSource_RedshiftServerlessStorageProperty struct {
	// The name of the Amazon Redshift Serverless workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftserverlessstorage.html#cfn-datazone-datasource-redshiftserverlessstorage-workgroupname
	//
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
}


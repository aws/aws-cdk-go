package previewawsiotanalyticsmixins


// The datastore activity that specifies where to store the processed data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datastoreProperty := &DatastoreProperty{
//   	DatastoreName: jsii.String("datastoreName"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-datastore.html
//
type CfnPipelinePropsMixin_DatastoreProperty struct {
	// The name of the data store where processed messages are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-datastore.html#cfn-iotanalytics-pipeline-datastore-datastorename
	//
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The name of the datastore activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-datastore.html#cfn-iotanalytics-pipeline-datastore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


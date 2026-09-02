package awsdataexchange


// Information about the published revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   revisionPublishedProperty := &RevisionPublishedProperty{
//   	DataSetId: jsii.String("dataSetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-revisionpublished.html
//
type CfnEventActionPropsMixin_RevisionPublishedProperty struct {
	// The data set ID of the published revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-revisionpublished.html#cfn-dataexchange-eventaction-revisionpublished-datasetid
	//
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
}


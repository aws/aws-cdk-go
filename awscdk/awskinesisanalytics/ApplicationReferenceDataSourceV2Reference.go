package awskinesisanalytics


// A reference to a ApplicationReferenceDataSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReferenceDataSourceV2Reference := &ApplicationReferenceDataSourceV2Reference{
//   	ApplicationReferenceDataSourceId: jsii.String("applicationReferenceDataSourceId"),
//   }
//
type ApplicationReferenceDataSourceV2Reference struct {
	// The Id of the ApplicationReferenceDataSource resource.
	ApplicationReferenceDataSourceId *string `field:"required" json:"applicationReferenceDataSourceId" yaml:"applicationReferenceDataSourceId"`
}


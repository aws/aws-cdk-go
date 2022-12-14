package awsroute53


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cidrRoutingConfigProperty := &cidrRoutingConfigProperty{
//   	collectionId: jsii.String("collectionId"),
//   	locationName: jsii.String("locationName"),
//   }
//
type CfnRecordSetGroup_CidrRoutingConfigProperty struct {
	// `CfnRecordSetGroup.CidrRoutingConfigProperty.CollectionId`.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// `CfnRecordSetGroup.CidrRoutingConfigProperty.LocationName`.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}


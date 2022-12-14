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
type CfnRecordSet_CidrRoutingConfigProperty struct {
	// `CfnRecordSet.CidrRoutingConfigProperty.CollectionId`.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// `CfnRecordSet.CidrRoutingConfigProperty.LocationName`.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}


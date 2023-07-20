package awsroute53


// The object that is specified in resource record set object when you are linking a resource record set to a CIDR location.
//
// A `LocationName` with an asterisk “*” can be used to create a default CIDR record. `CollectionId` is still required for default record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cidrRoutingConfigProperty := &CidrRoutingConfigProperty{
//   	CollectionId: jsii.String("collectionId"),
//   	LocationName: jsii.String("locationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-cidrroutingconfig.html
//
type CfnRecordSetGroup_CidrRoutingConfigProperty struct {
	// The CIDR collection ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-cidrroutingconfig.html#cfn-route53-recordsetgroup-cidrroutingconfig-collectionid
	//
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The CIDR collection location name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-cidrroutingconfig.html#cfn-route53-recordsetgroup-cidrroutingconfig-locationname
	//
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}


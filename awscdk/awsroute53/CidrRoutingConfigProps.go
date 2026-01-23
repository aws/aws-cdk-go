package awsroute53


// Properties for configuring CIDR routing in Route 53 resource record set objects.
//
// Example:
//   var myZone HostedZone
//
//
//   cidrCollection := route53.NewCfnCidrCollection(this, jsii.String("CidrCollection"), &CfnCidrCollectionProps{
//   	Name: jsii.String("test-collection"),
//   	Locations: []interface{}{
//   		&LocationProperty{
//   			CidrList: []*string{
//   				jsii.String("192.168.1.0/24"),
//   			},
//   			LocationName: jsii.String("my_location"),
//   		},
//   	},
//   })
//
//   route53.NewARecord(this, jsii.String("CidrRoutingConfig"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
//   	SetIdentifier: jsii.String("test"),
//   	CidrRoutingConfig: route53.CidrRoutingConfig_Create(&CidrRoutingConfigProps{
//   		CollectionId: cidrCollection.attrId,
//   		LocationName: jsii.String("test_location"),
//   	}),
//   })
//
type CidrRoutingConfigProps struct {
	// The CIDR collection ID.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The CIDR collection location name.
	// Default: `*`.
	//
	LocationName *string `field:"optional" json:"locationName" yaml:"locationName"`
}


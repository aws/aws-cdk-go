package awsroute53


// Properties for defining a `CfnCidrCollection`.
//
// Example:
//   var myZone hostedZone
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
//   		CollectionId: cidrCollection.AttrId,
//   		LocationName: jsii.String("test_location"),
//   	}),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html
//
type CfnCidrCollectionProps struct {
	// The name of a CIDR collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html#cfn-route53-cidrcollection-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A complex type that contains information about the list of CIDR locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-cidrcollection.html#cfn-route53-cidrcollection-locations
	//
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
}


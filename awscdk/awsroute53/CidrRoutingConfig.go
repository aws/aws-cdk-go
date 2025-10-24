package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for CIDR routing in Route 53 resource record set objects.
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
//   		CollectionId: cidrCollection.AttrId,
//   		LocationName: jsii.String("test_location"),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordset.html#cfn-route53-recordset-cidrroutingconfig
//
type CidrRoutingConfig interface {
	// The CIDR collection ID.
	CollectionId() *string
	// The CIDR collection location name.
	LocationName() *string
}

// The jsii proxy struct for CidrRoutingConfig
type jsiiProxy_CidrRoutingConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_CidrRoutingConfig) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CidrRoutingConfig) LocationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationName",
		&returns,
	)
	return returns
}


// Creates a new instance of CidrRoutingConfig.
func CidrRoutingConfig_Create(props *CidrRoutingConfigProps) CidrRoutingConfig {
	_init_.Initialize()

	if err := validateCidrRoutingConfig_CreateParameters(props); err != nil {
		panic(err)
	}
	var returns CidrRoutingConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CidrRoutingConfig",
		"create",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new instance of CidrRoutingConfig for default CIDR record.
//
// This method defines the locationName as `*`.
//
// Returns: A new instance of CidrRoutingConfig with the default location name as `*`.
func CidrRoutingConfig_WithDefaultLocationName(collectionId *string) CidrRoutingConfig {
	_init_.Initialize()

	if err := validateCidrRoutingConfig_WithDefaultLocationNameParameters(collectionId); err != nil {
		panic(err)
	}
	var returns CidrRoutingConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CidrRoutingConfig",
		"withDefaultLocationName",
		[]interface{}{collectionId},
		&returns,
	)

	return returns
}


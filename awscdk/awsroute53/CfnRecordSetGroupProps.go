package awsroute53


// Properties for defining a `CfnRecordSetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordSetGroupProps := &CfnRecordSetGroupProps{
//   	Comment: jsii.String("comment"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	HostedZoneName: jsii.String("hostedZoneName"),
//   	RecordSets: []interface{}{
//   		&RecordSetProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			AliasTarget: &AliasTargetProperty{
//   				DnsName: jsii.String("dnsName"),
//   				HostedZoneId: jsii.String("hostedZoneId"),
//
//   				// the properties below are optional
//   				EvaluateTargetHealth: jsii.Boolean(false),
//   			},
//   			CidrRoutingConfig: &CidrRoutingConfigProperty{
//   				CollectionId: jsii.String("collectionId"),
//   				LocationName: jsii.String("locationName"),
//   			},
//   			Failover: jsii.String("failover"),
//   			GeoLocation: &GeoLocationProperty{
//   				ContinentCode: jsii.String("continentCode"),
//   				CountryCode: jsii.String("countryCode"),
//   				SubdivisionCode: jsii.String("subdivisionCode"),
//   			},
//   			GeoProximityLocation: &GeoProximityLocationProperty{
//   				AwsRegion: jsii.String("awsRegion"),
//   				Bias: jsii.Number(123),
//   				Coordinates: &CoordinatesProperty{
//   					Latitude: jsii.String("latitude"),
//   					Longitude: jsii.String("longitude"),
//   				},
//   				LocalZoneGroup: jsii.String("localZoneGroup"),
//   			},
//   			HealthCheckId: jsii.String("healthCheckId"),
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   			HostedZoneName: jsii.String("hostedZoneName"),
//   			MultiValueAnswer: jsii.Boolean(false),
//   			Region: jsii.String("region"),
//   			ResourceRecords: []*string{
//   				jsii.String("resourceRecords"),
//   			},
//   			SetIdentifier: jsii.String("setIdentifier"),
//   			Ttl: jsii.String("ttl"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
//
type CfnRecordSetGroupProps struct {
	// *Optional:* Any comments you want to include about a change batch request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The ID of the hosted zone that you want to create records in.
	//
	// Specify either `HostedZoneName` or `HostedZoneId` , but not both. If you have multiple hosted zones with the same domain name, you must specify the hosted zone using `HostedZoneId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The name of the hosted zone that you want to create records in.
	//
	// You must include a trailing dot (for example, `www.example.com.` ) as part of the `HostedZoneName` .
	//
	// When you create a stack using an `AWS::Route53::RecordSet` that specifies `HostedZoneName` , AWS CloudFormation attempts to find a hosted zone whose name matches the `HostedZoneName` . If AWS CloudFormation can't find a hosted zone with a matching domain name, or if there is more than one hosted zone with the specified domain name, AWS CloudFormation will not create the stack.
	//
	// Specify either `HostedZoneName` or `HostedZoneId` , but not both. If you have multiple hosted zones with the same domain name, you must specify the hosted zone using `HostedZoneId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzonename
	//
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// A complex type that contains one `RecordSet` element for each record that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-recordsets
	//
	RecordSets interface{} `field:"optional" json:"recordSets" yaml:"recordSets"`
}


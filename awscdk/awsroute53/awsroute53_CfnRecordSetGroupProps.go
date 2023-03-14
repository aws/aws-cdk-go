package awsroute53


// Properties for defining a `CfnRecordSetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordSetGroupProps := &cfnRecordSetGroupProps{
//   	comment: jsii.String("comment"),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	hostedZoneName: jsii.String("hostedZoneName"),
//   	recordSets: []interface{}{
//   		&recordSetProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			aliasTarget: &aliasTargetProperty{
//   				dnsName: jsii.String("dnsName"),
//   				hostedZoneId: jsii.String("hostedZoneId"),
//
//   				// the properties below are optional
//   				evaluateTargetHealth: jsii.Boolean(false),
//   			},
//   			cidrRoutingConfig: &cidrRoutingConfigProperty{
//   				collectionId: jsii.String("collectionId"),
//   				locationName: jsii.String("locationName"),
//   			},
//   			failover: jsii.String("failover"),
//   			geoLocation: &geoLocationProperty{
//   				continentCode: jsii.String("continentCode"),
//   				countryCode: jsii.String("countryCode"),
//   				subdivisionCode: jsii.String("subdivisionCode"),
//   			},
//   			healthCheckId: jsii.String("healthCheckId"),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			multiValueAnswer: jsii.Boolean(false),
//   			region: jsii.String("region"),
//   			resourceRecords: []*string{
//   				jsii.String("resourceRecords"),
//   			},
//   			setIdentifier: jsii.String("setIdentifier"),
//   			ttl: jsii.String("ttl"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnRecordSetGroupProps struct {
	// *Optional:* Any comments you want to include about a change batch request.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The ID of the hosted zone that you want to create records in.
	//
	// Specify either `HostedZoneName` or `HostedZoneId` , but not both. If you have multiple hosted zones with the same domain name, you must specify the hosted zone using `HostedZoneId` .
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The name of the hosted zone that you want to create records in.
	//
	// You must include a trailing dot (for example, `www.example.com.` ) as part of the `HostedZoneName` .
	//
	// When you create a stack using an `AWS::Route53::RecordSet` that specifies `HostedZoneName` , AWS CloudFormation attempts to find a hosted zone whose name matches the `HostedZoneName` . If AWS CloudFormation can't find a hosted zone with a matching domain name, or if there is more than one hosted zone with the specified domain name, AWS CloudFormation will not create the stack.
	//
	// Specify either `HostedZoneName` or `HostedZoneId` , but not both. If you have multiple hosted zones with the same domain name, you must specify the hosted zone using `HostedZoneId` .
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// A complex type that contains one `RecordSet` element for each record that you want to create.
	RecordSets interface{} `field:"optional" json:"recordSets" yaml:"recordSets"`
}


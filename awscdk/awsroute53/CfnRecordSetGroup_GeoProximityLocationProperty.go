package awsroute53


// (Resource record sets only): A complex type that lets you specify where your resources are located.
//
// Only one of `LocalZoneGroup` , `Coordinates` , or `AWS Region` is allowed per request at a time.
//
// For more information about geoproximity routing, see [Geoproximity routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html) in the *Amazon Route 53 Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoProximityLocationProperty := &GeoProximityLocationProperty{
//   	AwsRegion: jsii.String("awsRegion"),
//   	Bias: jsii.Number(123),
//   	Coordinates: &CoordinatesProperty{
//   		Latitude: jsii.String("latitude"),
//   		Longitude: jsii.String("longitude"),
//   	},
//   	LocalZoneGroup: jsii.String("localZoneGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-geoproximitylocation.html
//
type CfnRecordSetGroup_GeoProximityLocationProperty struct {
	// The AWS Region the resource you are directing DNS traffic to, is in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-geoproximitylocation.html#cfn-route53-recordsetgroup-geoproximitylocation-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The bias increases or decreases the size of the geographic region from which Route 53 routes traffic to a resource.
	//
	// To use `Bias` to change the size of the geographic region, specify the applicable value for the bias:
	//
	// - To expand the size of the geographic region from which Route 53 routes traffic to a resource, specify a positive integer from 1 to 99 for the bias. Route 53 shrinks the size of adjacent regions.
	// - To shrink the size of the geographic region from which Route 53 routes traffic to a resource, specify a negative bias of -1 to -99. Route 53 expands the size of adjacent regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-geoproximitylocation.html#cfn-route53-recordsetgroup-geoproximitylocation-bias
	//
	Bias *float64 `field:"optional" json:"bias" yaml:"bias"`
	// Contains the longitude and latitude for a geographic region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-geoproximitylocation.html#cfn-route53-recordsetgroup-geoproximitylocation-coordinates
	//
	Coordinates interface{} `field:"optional" json:"coordinates" yaml:"coordinates"`
	// Specifies an AWS Local Zone Group.
	//
	// A local Zone Group is usually the Local Zone code without the ending character. For example, if the Local Zone is `us-east-1-bue-1a` the Local Zone Group is `us-east-1-bue-1` .
	//
	// You can identify the Local Zones Group for a specific Local Zone by using the [describe-availability-zones](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-availability-zones.html) CLI command:
	//
	// This command returns: `"GroupName": "us-west-2-den-1"` , specifying that the Local Zone `us-west-2-den-1a` belongs to the Local Zone Group `us-west-2-den-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup-geoproximitylocation.html#cfn-route53-recordsetgroup-geoproximitylocation-localzonegroup
	//
	LocalZoneGroup *string `field:"optional" json:"localZoneGroup" yaml:"localZoneGroup"`
}


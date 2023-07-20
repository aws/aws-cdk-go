package awsroute53


// A complex type that contains information about a tag that you want to add or edit for the specified health check or hosted zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedZoneTagProperty := &HostedZoneTagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html
//
type CfnHostedZone_HostedZoneTagProperty struct {
	// The value of `Key` depends on the operation that you want to perform:.
	//
	// - *Add a tag to a health check or hosted zone* : `Key` is the name that you want to give the new tag.
	// - *Edit a tag* : `Key` is the name of the tag that you want to change the `Value` for.
	// - *Delete a key* : `Key` is the name of the tag you want to remove.
	// - *Give a name to a health check* : Edit the default `Name` tag. In the Amazon Route 53 console, the list of your health checks includes a *Name* column that lets you see the name that you've given to each health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html#cfn-route53-hostedzone-hostedzonetag-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of `Value` depends on the operation that you want to perform:.
	//
	// - *Add a tag to a health check or hosted zone* : `Value` is the value that you want to give the new tag.
	// - *Edit a tag* : `Value` is the new value that you want to assign the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html#cfn-route53-hostedzone-hostedzonetag-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


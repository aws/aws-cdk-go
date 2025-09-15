package awsconfig


// A reference to a AggregationAuthorization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationAuthorizationReference := &AggregationAuthorizationReference{
//   	AggregationAuthorizationArn: jsii.String("aggregationAuthorizationArn"),
//   	AuthorizedAccountId: jsii.String("authorizedAccountId"),
//   	AuthorizedAwsRegion: jsii.String("authorizedAwsRegion"),
//   }
//
type AggregationAuthorizationReference struct {
	// The ARN of the AggregationAuthorization resource.
	AggregationAuthorizationArn *string `field:"required" json:"aggregationAuthorizationArn" yaml:"aggregationAuthorizationArn"`
	// The AuthorizedAccountId of the AggregationAuthorization resource.
	AuthorizedAccountId *string `field:"required" json:"authorizedAccountId" yaml:"authorizedAccountId"`
	// The AuthorizedAwsRegion of the AggregationAuthorization resource.
	AuthorizedAwsRegion *string `field:"required" json:"authorizedAwsRegion" yaml:"authorizedAwsRegion"`
}


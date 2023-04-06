// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// custom domain name attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   domainNameAttributes := &DomainNameAttributes{
//   	Name: jsii.String("name"),
//   	RegionalDomainName: jsii.String("regionalDomainName"),
//   	RegionalHostedZoneId: jsii.String("regionalHostedZoneId"),
//   }
//
// Experimental.
type DomainNameAttributes struct {
	// domain name string.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain name associated with the regional endpoint for this custom domain name.
	// Experimental.
	RegionalDomainName *string `field:"required" json:"regionalDomainName" yaml:"regionalDomainName"`
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// Experimental.
	RegionalHostedZoneId *string `field:"required" json:"regionalHostedZoneId" yaml:"regionalHostedZoneId"`
}


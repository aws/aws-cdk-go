package awsroute53profiles


// Properties for defining a `CfnProfileResourceAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileResourceAssociationProps := &CfnProfileResourceAssociationProps{
//   	Name: jsii.String("name"),
//   	ProfileId: jsii.String("profileId"),
//   	ResourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	ResourceProperties: jsii.String("resourceProperties"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html
//
type CfnProfileResourceAssociationProps struct {
	// Name of the Profile resource association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Profile ID of the Profile that the resources are associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The Amazon Resource Name (ARN) of the resource association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// If the DNS resource is a DNS Firewall rule group, this indicates the priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-resourceproperties
	//
	ResourceProperties *string `field:"optional" json:"resourceProperties" yaml:"resourceProperties"`
}


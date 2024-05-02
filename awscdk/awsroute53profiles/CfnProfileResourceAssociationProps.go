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
	// The name of an association between the  Profile and resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The arn of the resource that you associated to the  Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileresourceassociation.html#cfn-route53profiles-profileresourceassociation-resourceproperties
	//
	ResourceProperties *string `field:"optional" json:"resourceProperties" yaml:"resourceProperties"`
}


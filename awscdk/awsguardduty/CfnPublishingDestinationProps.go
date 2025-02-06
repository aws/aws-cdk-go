package awsguardduty


// Properties for defining a `CfnPublishingDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPublishingDestinationProps := &CfnPublishingDestinationProps{
//   	DestinationProperties: &CFNDestinationPropertiesProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	DestinationType: jsii.String("destinationType"),
//   	DetectorId: jsii.String("detectorId"),
//
//   	// the properties below are optional
//   	Tags: []tagItemProperty{
//   		&tagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-publishingdestination.html
//
type CfnPublishingDestinationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-publishingdestination.html#cfn-guardduty-publishingdestination-destinationproperties
	//
	DestinationProperties interface{} `field:"required" json:"destinationProperties" yaml:"destinationProperties"`
	// The type of resource for the publishing destination.
	//
	// Currently only Amazon S3 buckets are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-publishingdestination.html#cfn-guardduty-publishingdestination-destinationtype
	//
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// The ID of the GuardDuty detector associated with the publishing destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-publishingdestination.html#cfn-guardduty-publishingdestination-detectorid
	//
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-publishingdestination.html#cfn-guardduty-publishingdestination-tags
	//
	Tags *[]*CfnPublishingDestination_TagItemProperty `field:"optional" json:"tags" yaml:"tags"`
}


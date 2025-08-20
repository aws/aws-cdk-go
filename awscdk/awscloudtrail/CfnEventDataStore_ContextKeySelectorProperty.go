package awscloudtrail


// An object that contains information types to be included in CloudTrail enriched events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextKeySelectorProperty := &ContextKeySelectorProperty{
//   	EqualTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-contextkeyselector.html
//
type CfnEventDataStore_ContextKeySelectorProperty struct {
	// A list of keys defined by Type to be included in CloudTrail enriched events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-contextkeyselector.html#cfn-cloudtrail-eventdatastore-contextkeyselector-equals
	//
	EqualTo *[]*string `field:"required" json:"equalTo" yaml:"equalTo"`
	// Specifies the type of the event record field in ContextKeySelector.
	//
	// Valid values include RequestContext, TagContext.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-contextkeyselector.html#cfn-cloudtrail-eventdatastore-contextkeyselector-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}


package awsmediaconnectalpha


// Options for Entitlement.
//
// Example:
//   var stack Stack
//
//
//   // Import an entitlement from another AWS account
//   entitlement := awsmediaconnectalpha.FlowEntitlement_FromFlowEntitlementArn(stack, jsii.String("ImportedEntitlement"), jsii.String("arn:aws:mediaconnect:us-west-2:111122223333:entitlement:1-11111111111111111111111111111111:MyEntitlement"))
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Entitlement(&EntitlementSource{
//   		Entitlement: entitlement,
//   	}),
//   })
//
// Experimental.
type EntitlementSource struct {
	// The entitlement that allows you to subscribe to content that comes from another AWS account.
	// Experimental.
	Entitlement IFlowEntitlement `field:"required" json:"entitlement" yaml:"entitlement"`
	// Options to decrypt incoming feed.
	// Default: - no decryption.
	//
	// Experimental.
	Decryption *StaticKeyEncryption `field:"optional" json:"decryption" yaml:"decryption"`
}


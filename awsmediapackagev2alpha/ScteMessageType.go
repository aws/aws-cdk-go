package awsmediapackagev2alpha


// SCTE-35 message type options available.
// Experimental.
type ScteMessageType string

const (
	// Option for SPLICE_INSERT.
	// Experimental.
	ScteMessageType_SPLICE_INSERT ScteMessageType = "SPLICE_INSERT"
	// Option for BREAK.
	// Experimental.
	ScteMessageType_BREAK ScteMessageType = "BREAK"
	// Option for PROVIDER_ADVERTISEMENT.
	// Experimental.
	ScteMessageType_PROVIDER_ADVERTISEMENT ScteMessageType = "PROVIDER_ADVERTISEMENT"
	// Option for DISTRIBUTOR_ADVERTISEMENT.
	// Experimental.
	ScteMessageType_DISTRIBUTOR_ADVERTISEMENT ScteMessageType = "DISTRIBUTOR_ADVERTISEMENT"
	// Option for PROVIDER_PLACEMENT_OPPORTUNITY.
	// Experimental.
	ScteMessageType_PROVIDER_PLACEMENT_OPPORTUNITY ScteMessageType = "PROVIDER_PLACEMENT_OPPORTUNITY"
	// Option for DISTRIBUTOR_PLACEMENT_OPPORTUNITY.
	// Experimental.
	ScteMessageType_DISTRIBUTOR_PLACEMENT_OPPORTUNITY ScteMessageType = "DISTRIBUTOR_PLACEMENT_OPPORTUNITY"
	// Option for PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY.
	// Experimental.
	ScteMessageType_PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY ScteMessageType = "PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY"
	// Option for DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY.
	// Experimental.
	ScteMessageType_DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY ScteMessageType = "DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY"
	// Option for PROGRAM.
	// Experimental.
	ScteMessageType_PROGRAM ScteMessageType = "PROGRAM"
)


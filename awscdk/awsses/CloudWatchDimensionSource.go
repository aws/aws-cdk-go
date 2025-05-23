package awsses


// Source for CloudWatch dimension.
type CloudWatchDimensionSource string

const (
	// Amazon SES retrieves the dimension name and value from a header in the email.
	//
	// Note: You can't use any of the following email headers as the Dimension Name:
	// `Received`, `To`, `From`, `DKIM-Signature`, `CC`, `message-id`, or `Return-Path`.
	CloudWatchDimensionSource_EMAIL_HEADER CloudWatchDimensionSource = "EMAIL_HEADER"
	// Amazon SES retrieves the dimension name and value from a tag that you specified in a link.
	// See: https://docs.aws.amazon.com/ses/latest/dg/faqs-metrics.html#sending-metric-faqs-clicks-q5
	//
	CloudWatchDimensionSource_LINK_TAG CloudWatchDimensionSource = "LINK_TAG"
	// Amazon SES retrieves the dimension name and value from a tag that you specify by using the `X-SES-MESSAGE-TAGS` header or the Tags API parameter.
	//
	// You can also use the Message Tag value source to create dimensions based on Amazon SES auto-tags.
	// To use an auto-tag, type the complete name of the auto-tag as the Dimension Name. For example,
	// to create a dimension based on the configuration set auto-tag, use `ses:configuration-set` for the
	// Dimension Name, and the name of the configuration set for the Default Value.
	// See: https://docs.aws.amazon.com/ses/latest/dg/monitor-using-event-publishing.html#event-publishing-how-works
	//
	CloudWatchDimensionSource_MESSAGE_TAG CloudWatchDimensionSource = "MESSAGE_TAG"
)


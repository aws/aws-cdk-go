package awsses


// Email sending event.
type EmailSendingEvent string

const (
	// The send request was successful and SES will attempt to deliver the message to the recipient's mail server.
	//
	// (If account-level or global suppression is
	// being used, SES will still count it as a send, but delivery is suppressed.)
	EmailSendingEvent_SEND EmailSendingEvent = "SEND"
	// SES accepted the email, but determined that it contained a virus and didn’t attempt to deliver it to the recipient’s mail server.
	EmailSendingEvent_REJECT EmailSendingEvent = "REJECT"
	// (Hard bounce) The recipient's mail server permanently rejected the email.
	//
	// (Soft bounces are only included when SES fails to deliver the email after
	// retrying for a period of time.)
	EmailSendingEvent_BOUNCE EmailSendingEvent = "BOUNCE"
	// The email was successfully delivered to the recipient’s mail server, but the recipient marked it as spam.
	EmailSendingEvent_COMPLAINT EmailSendingEvent = "COMPLAINT"
	// SES successfully delivered the email to the recipient's mail server.
	EmailSendingEvent_DELIVERY EmailSendingEvent = "DELIVERY"
	// The recipient received the message and opened it in their email client.
	EmailSendingEvent_OPEN EmailSendingEvent = "OPEN"
	// The recipient clicked one or more links in the email.
	EmailSendingEvent_CLICK EmailSendingEvent = "CLICK"
	// The email wasn't sent because of a template rendering issue.
	//
	// This event type
	// can occur when template data is missing, or when there is a mismatch between
	// template parameters and data. (This event type only occurs when you send email
	// using the `SendTemplatedEmail` or `SendBulkTemplatedEmail` API operations.)
	EmailSendingEvent_RENDERING_FAILURE EmailSendingEvent = "RENDERING_FAILURE"
	// The email couldn't be delivered to the recipient’s mail server because a temporary issue occurred.
	//
	// Delivery delays can occur, for example, when the recipient's inbox
	// is full, or when the receiving email server experiences a transient issue.
	EmailSendingEvent_DELIVERY_DELAY EmailSendingEvent = "DELIVERY_DELAY"
	// The email was successfully delivered, but the recipient updated their subscription preferences by clicking on an unsubscribe link as part of your subscription management.
	EmailSendingEvent_SUBSCRIPTION EmailSendingEvent = "SUBSCRIPTION"
)


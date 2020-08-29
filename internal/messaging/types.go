package messaging

/**
*	All necessary types for building response through Messenger API webhook
*	It strictly respect `messages` event standards for attributes naming and order
 */

// User : User ID struct
type User struct {
	ID string `json:"id,omitempty"`
}

// Messaging struct
type Messaging struct {
	Sender    User    `json:"sender,omitempty"`
	Recipient User    `json:"recipient,omitempty"`
	Timestamp int     `json:"timestamp,omitempty"`
	Message   Message `json:"message,omitempty"`
}

// Message struct
type Message struct {
	MID        string `json:"mid,omitempty"`
	Text       string `json:"text,omitempty"`
	QuickReply *struct {
		Payload string `json:"payload,omitempty"`
	} `json:"quick_reply,omitempty"`
	Attachments *[]Attachment `json:"attachments,omitempty"`
	Attachment  *Attachment   `json:"attachment,omitempty"`
}

// Attachment struct
type Attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

// Response struct
type Response struct {
	Recipient User    `json:"recipient,omitempty"`
	Message   Message `json:"message,omitempty"`
}

// Payload struct
type Payload struct {
	URL string `json:"url,omitempty"`
}

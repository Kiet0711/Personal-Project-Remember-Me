// Package channels defines the NotificationChannel interface and per-channel
// implementations: zalo, messenger, desktop.
//
// Zalo is the preferred channel; Messenger and Desktop are secondary.
package channels

// Channel is the interface every notification channel must satisfy.
type Channel interface {
	// Name returns a short identifier for the channel (e.g. "zalo", "messenger", "desktop").
	Name() string

	// Send delivers a payload to the configured user/recipient.
	// Implementation: to be added when the channel is wired (Phase 11+).
	Send(userID int64, payload []byte) error
}

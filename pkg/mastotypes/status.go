/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package mastotypes

// StatusRequest represents a mastodon-api status POST request, as defined here: https://docs.joinmastodon.org/methods/statuses/
// It should be used at the path https://mastodon.example/api/v1/statuses
type StatusRequest struct {
	// Text content of the status. If media_ids is provided, this becomes optional. Attaching a poll is optional while status is provided.
	Status string `form:"status"`
	// Array of Attachment ids to be attached as media. If provided, status becomes optional, and poll cannot be used.
	MediaIDs []string `form:"media_ids"`
	// Poll to include with this status.
	Poll *PollRequest `form:"poll"`
	// ID of the status being replied to, if status is a reply
	InReplyToID string `form:"in_reply_to_id"`
	// Mark status and attached media as sensitive?
	Sensitive bool `form:"sensitive"`
	// Text to be shown as a warning or subject before the actual content. Statuses are generally collapsed behind this field.
	SpoilerText string `form:"spoiler_text"`
	// Visibility of the posted status. Enumerable oneOf public, unlisted, private, direct.
	Visibility string `form:"visibility"`
	// ISO 8601 Datetime at which to schedule a status. Providing this paramter will cause ScheduledStatus to be returned instead of Status. Must be at least 5 minutes in the future.
	ScheduledAt string `form:"scheduled_at"`
	// ISO 639 language code for this status.
	Language string `form:"language"`
}

// Status represents a mastodon-api Status type, as defined here: https://docs.joinmastodon.org/entities/status/
type Status struct {
	// ID of the status in the database.
	ID string `json:"id"`
	// The date when this status was created (ISO 8601 Datetime)
	CreatedAt string `json:"created_at"`
	// ID of the status being replied.
	InReplyToID string `json:"in_reply_to_id"`
	// ID of the account being replied to.
	InReplyToAccountID string `json:"in_reply_to_account_id"`
	// Is this status marked as sensitive content?
	Sensitive bool `json:"sensitive"`
	// Subject or summary line, below which status content is collapsed until expanded.
	SpoilerText string `json:"spoiler_text"`
	// Visibility of this status.
	// 	public = Visible to everyone, shown in public timelines.
	// 	unlisted = Visible to public, but not included in public timelines.
	// 	private = Visible to followers only, and to any mentioned users.
	// 	direct = Visible only to mentioned users.
	Visibility string `json:"visibility"`
	// Primary language of this status. (ISO 639 Part 1 two-letter language code)
	Language string `json:"language"`
	// URI of the status used for federation.
	URI string `json:"uri"`
	// A link to the status's HTML representation.
	URL string `json:"url"`
	// How many replies this status has received.
	RepliesCount int `json:"replies_count"`
	// How many boosts this status has received.
	ReblogsCount int `json:"reblogs_count"`
	// How many favourites this status has received.
	FavouritesCount int `json:"favourites_count"`
	// Have you favourited this status?
	Favourited bool `json:"favourited"`
	// Have you boosted this status?
	Reblogged bool `json:"reblogged"`
	// Have you muted notifications for this status's conversation?
	Muted bool `json:"muted"`
	// Have you bookmarked this status?
	Bookmarked bool `json:"bookmarked"`
	// Have you pinned this status? Only appears if the status is pinnable.
	Pinned bool `json:"pinned"`
	// HTML-encoded status content.
	Content string `json:"content"`
	// The status being reblogged.
	Reblog *Status `json:"reblog"`
	// The application used to post this status.
	Application *Application `json:"application"`
	// The account that authored this status.
	Account *Account `json:"account"`
	// Media that is attached to this status.
	MediaAttachments []Attachment `json:"media_attachments"`
	// Mentions of users within the status content.
	Mentions []Mention `json:"mentions"`
	// Hashtags used within the status content.
	Tags []Tag `json:"tags"`
	// Custom emoji to be used when rendering status content.
	Emojis []Emoji `json:"emojis"`
	// Preview card for links included within status content.
	Card *Card `json:"card"`
	// The poll attached to the status.
	Poll *Poll `json:"poll"`
	// Plain-text source of a status. Returned instead of content when status is deleted,
	// so the user may redraft from the source text without the client having to reverse-engineer
	// the original text from the HTML content.
	Text string `json:"text"`
}

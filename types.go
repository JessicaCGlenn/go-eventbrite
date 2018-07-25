package eventbrite

import (
	"bytes"
	"fmt"
	"time"
)

// When an error occurs during an API request, you’ll get a response with an error HTTP status
// (in the 400 or 500 range), as well as a JSON response containing more information about the error.
//
// https://www.eventbrite.co.uk/developer/v3/api_overview/errors/#ebapi-errors
type Error struct {
	// The error key contains a constant string value for error - in this case, VENUE_AND_ONLINE - and
	// is what you should key your error handling off of, as this string won’t change depending on locale
	// or as we change the API over time
	Err string `json:"error,omitempty"`
	// The error_description key is for developer information only and will usually contain a more informative
	// explanation for the error, should you be confused. You should not display this string to your users;
	// it’s often very technical and may not be localized to their language
	Description string `json:"error_description,omitempty"`
	// The status_code value just mirrors the HTTP status code you got as part of the request. It’s there as
	// a convenience if your HTTP library makes it very hard to get status codes, or has one error handler
	// for all error codes
	Status int `json:"status_code,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Eventbrite API: [Status code - %d] %s", e.Status, e.Description)
}

// The ISO 3166 alpha-2 code of a country.
type CountryCode string

// An ISO 4217 3-character code of a currency
type CurrencyCode string

type Currency struct {
	Currency CurrencyCode `json:"currency,omitempty"`
	Value    float32      `json:"value,omitempty"`
	Display  string       `json:"display,omitempty"`
}

type Date struct {
	Time time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte("\""), []byte(""), -1)
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		fmt.Println(err)
	}

	d.Time = t
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Time.Format("2006-01-02") + "\""), nil
}

type DateTime struct {
	Time time.Time
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte("\""), []byte(""), -1)
	t, err := time.Parse("2006-01-02T15:04:05Z", string(data))
	if err != nil {
		fmt.Println(err)
	}

	d.Time = t
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Time.Format("2006-01-02T15:04:05Z") + "\""), nil
}

// Timezone is an object with details about a timezone
type Timezone struct {
	// Timezone id
	ID string `json:"id,omitempty"`
	// The timezone identifier as defined by the IANA Time Zone Database
	Timezone string `json:"timezone,omitempty"`
	// The localized name for the timezone
	Label string `json:"label,omitempty"`
}

// All API endpoints which return multiple objects will return paginated responses; as well as the
// list of objects (which will usually be under a key like events or attendees, depending on the endpoint)
// there will also be a pagination key:
//
// see @https://www.eventbrite.com/developer/v3/api_overview/pagination/#ebapi-paginated-responses
type Pagination struct {
	ObjectCount  int  `json:"object_count,omitempty"`
	PageNumber   int  `json:"page_number,omitempty"`
	PageSize     int  `json:"page_size,omitempty"`
	PageCount    int  `json:"page_count,omitempty"`
	HasMoreItems bool `json:"has_more_items,omitempty"`
}

// Returned for fields which represent HTML, like event names and descriptions.
// The html key represents the original HTML (which _should_ be sanitized and free from injected script tags etc.,
// but as always, be careful what you put in your DOM), while the text key is a stripped version useful for places
// where you can’t or don’t need to display the full HTML version.
//
// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-multipart-text
type MultipartText struct {
	Text string `json:"text,omitempty"`
	Html string `json:"html,omitempty"`
}

// A combination of a timezone from the Olson specification as a string, and two datetime values, one for
// the UTC time represented and one for the local time in the named timezone.
//
// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-datetime-with-timezone
type DatetimeTz struct {
	Timezone string `json:"timezone,omitempty"`
	Utc      string `json:"utc,omitempty"`
	Local    string `json:"local,omitempty"`
}

// Country is an object with details about a country
//
// https://www.eventbrite.com/developer/v3/response_formats/system/#ebapi-countries
type Country struct {
	// The country identifier as defined by the ISO 3166 standard
	Code CountryCode `json:"code,omitempty"`
	// The readable name of the country
	Label string `json:"label,omitempty"`
}

// Region is an object with details about a region
//
// https://www.eventbrite.com/developer/v3/response_formats/system/#ebapi-region
type Region struct {
	// The associated country code to this region
	CountryCode string `json:"country_code,omitempty"`
	// The region identifier as defined by the ISO 3166 standard
	Code string `json:"code,omitempty"`
	// The readable name of the region
	Label string `json:"label,omitempty"`
}

// Image is an object with details about a given image.
//
// https://www.eventbrite.com/developer/v3/response_formats/image/#ebapi-image
type Image struct {
	// The image’s ID
	ID string `json:"id,omitempty"`
	// The URL of the image
	Url string `json:"url,omitempty"`
}

// A location where an event happens.
//
// https://www.eventbrite.com/developer/v3/response_formats/venue/#ebapi-venue
type Venue struct {
	// The value name
	Name string `json:"name,omitempty"`
	// The address of the venue
	Address Address `json:"address,omitempty"`
}

// Though address formatting varies considerably between different countries and regions, Eventbrite
// still has a common address return format to keep things consistent.
//
// https://www.eventbrite.com/developer/v3/response_formats/basic/#ebapi-address
type Address struct {
	// The street/location address (part 1)
	Address1 string `json:"address_1,omitempty"`
	// The street/location address (part 2)
	Address2 string `json:"address_2,omitempty"`
	// The city
	City string `json:"city,omitempty"`
	// The ISO 3166-2 2- or 3-character region code for the state, province, region, or district
	Region string `json:"region,omitempty"`
	// The postal code
	PostalCode string `json:"postal_code,omitempty"`
	// The ISO 3166-1 2-character international code for the country
	Country string `json:"country,omitempty"`
	// The latitude portion of the address coordinates
	Latitude string `json:"latitude,omitempty"`
	// The longitude portion of the address coordinates
	Longitude string `json:"longitude,omitempty"`
	// The format of the address display localized to the address country
	LocalizedAddressDisplay string `json:"localized_address_display,omitempty"`
	// The format of the address’s area display localized to the address country
	LocalizedAreaDisplay string `json:"localized_area_display,omitempty"`
	//     The multi-line format order of the address display localized to the address country, where each line is an item in the list
	LocalizedMultiLineAddressDisplay []interface{} `json:"localized_multi_line_address_display,omitempty"`
}

// A grouping entity that Eventbrite uses to display as the owner of events. Contains name
// and contact information.
//
// https://www.eventbrite.com/developer/v3/response_formats/organizer/#ebapi-std:format-organizer
type Organizer struct {
	// The organizer name
	Name string `json:"name,omitempty"`
	// The description of the organizer (may be very long and contain significant formatting)
	Description *MultipartText `json:"description,omitempty"`
	// The URL to the organizer’s page on Eventbrite
	Url string `json:"url,omitempty"`
}

// An overarching category that an event falls into (vertical). Examples are “Music”, and “Endurance”.
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-category
type Category struct {
	// Category ID
	ID string `json:"id,omitempty"`
	// he category name
	Name string `json:"name,omitempty"`
	// The category name localized to the current locale (if available)
	NameLocalized string `json:"name_localized,omitempty"`
	// A shorter name for display in sidebars and other small spaces.
	ShortName string `json:"short_name,omitempty"`
	// List of subcategories, only shown on some endpoints.
	ShortNameLocalized string `json:"short_name_localized,omitempty"`
	SubCategories      []SubCategory	`json:"sub_categories,omitempty"`
}

// A more specific category that an event falls into, sitting underneath a category.
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-subcategory
type SubCategory struct {
	// Subcategory ID
	ID string `json:"id,omitempty"`
	// The category name
	Name string `json:"name,omitempty"`
	// The category this belongs to
	ParentCategory *Category `json:"parent_category,omitempty"`
}

// https://www.eventbrite.com/developer/v3/endpoints/events/#ebapi-get-events-id-display-settings
type EventSettings struct {
	// Whether to display the start date on the event listing
	ShowStartDate bool `json:"display_settings.show_start_date,omitempty"`
	// Whether to display the end date on the event listing
	ShowEndDate bool `json:"display_settings.show_end_date,omitempty"`
	// Whether to display event start and end time on the event listing
	ShowStartEndTime bool `json:"display_settings.show_start_end_time,omitempty"`
	// Whether to display the event timezone on the event listing
	ShowTimezone bool `json:"display_settings.show_timezone,omitempty"`
	// Whether to display a map to the venue on the event listing
	ShowMap bool `json:"display_settings.show_map,omitempty"`
	// Whether to display the number of remaining tickets
	ShowRemaining bool `json:"display_settings.show_remaining,omitempty"`
	// Whether to display a link to the organizer’s Facebook profile
	ShowOrganizerFacebook bool `json:"display_settings.show_organizer_facebook,omitempty"`
	// Whether to display a link to the organizer’s Twitter profile
	ShowOrganizerTwitter bool `json:"display_settings.show_organizer_twitter,omitempty"`
	// Whether to display which of the user’s Facebook friends are going
	ShowFacebookFriendsGoing bool `json:"display_settings.show_facebook_friends_going,omitempty"`
	// Which terminology should be used to refer to the event (Valid choices are: tickets_vertical, or endurance_vertical)
	ShowAttendeeList bool `json:"display_settings.show_attendee_list,omitempty"`
}

// This is an object representing one of the possible ticket classes (types of ticket) for an event
//
// https://www.eventbrite.com/developer/v3/response_formats/event/#ebapi-ticket-class
type TicketClass struct {
	ID string `json:"id,omitempty"`
	// The ticket class’ name
	Name string `json:"name,omitempty"`
	// The ticket’s description. (optional)
	Description string `json:"description,omitempty"`
	// The display cost of the ticket (paid only)
	Cost Currency `json:"cost,omitempty"`
	// The display fee of the ticket (paid only)
	Fee Currency `json:"fee,omitempty"`
	// If the ticket is a donation
	Donation bool `json:"donation,omitempty"`
	// If the ticket is a free ticket
	Free bool `json:"free,omitempty"`
	// Minimum number that can be bought per order
	MinimumQuantity int `json:"minimum_quantity,omitempty"`
	// Maximum number that can be bought per order
	MaximumQuantity int `json:"maximum_quantity,omitempty"`
	// The event the ticket class is for
	EventID string `json:"event_id,omitempty"`
	// The event the ticket class is for
	Event Event `json:"event,omitempty"`
	// PRIVATE FIELDS
	// Only shown to people with event owner permission
	// How many of these tickets are available to be sold overall
	QuantityTotal int `json:"quantity_total,omitempty"`
	// How many of these tickets have already been sold and confirmed (does not include tickets being checked out right now)
	QuantitySold int `json:"quantity_sold,omitempty"`
	// If the ticket is hidden from the public
	Hidden bool `json:"hidden,omitempty"`
	// When sales for this ticket start
	SalesStart string `json:"sales_start,omitempty"`
	// When sales for this ticket end
	SalesEnd string `json:"sales_end,omitempty"`
	// The ID of another ticket class that, when it sells out, will trigger sales of this class to start
	SalesStartAfter string `json:"sales_start_after,omitempty"`
	// If the fee should be included in the displayed cost (cannot be set along with split_fee)
	IncludeFee bool `json:"include_fee,omitempty"`
	// If the payment fee should be included in the displayed cost and the eventbrite fee is shown separately
	SplitFee bool `json:"split_fee,omitempty"`
	// If the description should be hidden on the event page (will remove description from public responses too)
	HideDescription bool `json:"hide_description,omitempty"`
	// If the ticket should be hidden when not on sale
	AutoHide bool `json:"auto_hide,omitempty"`
	// Override the time at which auto hide disables itself to show the ticket (otherwise it’s sales_start)
	AutoHideBefore string `json:"auto_hide_before,omitempty"`
	// Override the time at which auto hide enables itself to re-hide the ticket (otherwise it’s sales_end)
	AutoHideAfter string `json:"auto_hide_after,omitempty"`
}

// An entity that Eventbrite uses to allow event organizer to utilize tracking pixels on their events
//
// https://www.eventbrite.com/developer/v3/response_formats/tracking_beacon/#ebapi-tracking-beacon
type TrackingBeacon struct {
	// The tracking beacon id
	ID string
	// The tracking beacon third party type. Allowed types are: Facebook Pixel,
	// Twitter Ads, AdWords, Google Analytics, Simple Image Pixel, Adroll iPixel
	TrackingType string
	// The id of the event where the tracking beacon will load your tracking pixel
	EventID string
	// The id of the user where the tracking beacon will load this tracking pixel on all of their events
	UserID string
	// The third party id that they have given you to fire on your event page
	PixelID string
	// The tracking pixel meta information that determines where your pixel will fire
	Triggers interface{}
}

// An object representing a single webhook associated with the account
type Webhook struct {
	// The url that the webhook will send data to when it is triggered
	EndpointUrl string `json:"endpoint_url,omitempty"`
	// One or any combination of actions that will cause this webhook to fire
	Actions string `json:"actions,omitempty"`
}

// Attendee is an object representing the details of one or more people coming to the event
// Attendee objects are considered private and are only available to the event owner
type Attendee struct {
	// When the attendee was created (order placed)
	Created DateTime `json:"created,omitempty"`
	// When the attendee was last changed
	Changed DateTime `json:"changed,omitempty"`
	// The name of the ticket_class at the time of registration
	TicketClassName string `json:"ticket_class_name,omitempty"`
	// The attendee’s basic profile information
	Profile *AttendeeProfile `json:"profile,omitempty"`
	// The attendee’s basic profile information
	Addresses *AttendeeAddresses `json:"addresses,omitempty"`
	// The attendee’s answers to any custom questions (optional)
	Answers *AttendeeAnswers `json:"answers,omitempty"`
	// The attendee’s entry barcode information
	Barcodes *AttendeeBarcodes `json:"barcodes,omitempty"`
	// The attendee’s team information (optional)
	Team *AttendeeTeam `json:"team,omitempty"`
	// The attendee’s affiliate code (optional)
	//
	// Not documented
	Affiliate interface{} `json:"affiliate,omitempty"`
	// If the attendee is checked in
	CheckedIn bool `json:"checked_in,omitempty"`
	// If the attendee is cancelled
	Cancelled bool `json:"cancelled,omitempty"`
	// If the attendee is refunded
	Refunded bool `json:"refunded,omitempty"`
	// The status of the attendee (scheduled to be deprecated)
	Status string `json:"status,omitempty"`
	// The event id that this attendee is attending
	EventID string `json:"event_id,omitempty"`
	// The event this attendee is attending
	Event *Event `json:"event,omitempty"`
	// The order id this attendee is part of
	OrderID string `json:"order_id,omitempty"`
	// The order this attendee is part of
	Order *Order `json:"order,omitempty"`
	// The guestlist id for this attendee. If this is null it means that this is not a guest
	GuestListID string `json:"guestlist_id,omitempty"`
	// The guest of for the guest. If this is null it means that this is not a guest
	InvitedBy string `json:"invited_by,omitempty"`
	// The promotional code applied to this attendee
	//
	// Not documented
	PromotionalCode interface{} `json:"promotional_code,omitempty"`
	// The bib number assigned to this attendee if one exists for a race or endurance event
	//
	// Not documented
	AssignedNumber interface{} `json:"assigned_number,omitempty"`
}

// Contains the attendee’s personal information
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-std:format-attendee-profile
type AttendeeProfile struct {
	// The attendee’s name. Use this in preference to first_name/last_name/etc. if possible for
	// forward compatibility with non-Western names
	Name string `json:"name,omitempty"`
	// The attendee’s email address
	Email string `json:"email,omitempty"`
	// The attendee’s first name
	FirstName string `json:"first_name,omitempty"`
	// The attendee’s last name
	LastName string `json:"last_name,omitempty"`
	// The title or honoraria used in front of the name (Mr., Mrs., etc.) (optional)
	Prefix string `json:"prefix,omitempty"`
	// The suffix at the end of the name (e.g. Jr, Sr) (optional)
	Suffix string `json:"suffix,omitempty"`
	// The attendee’s age (optional)
	Age int `json:"age,omitempty"`
	// The attendee’s job title (optional)
	JobTitle string `json:"job_title,omitempty"`
	// The attendee’s company name (optional)
	Company string `json:"company,omitempty"`
	// The attendee’s website address (optional)
	Website string `json:"website,omitempty"`
	// The attendee’s blog address (optional)
	Blog string `json:"blog,omitempty"`
	// The attendee’s gender (currently one of “male” or “female”) (optional)
	Gender string `json:"gender,omitempty"`
	// The attendee’s birth date (optional)
	BirthDate Date `json:"birth_date,omitempty"`
	// The attendee’s cell/mobile phone number, as formatted by them (optional)
	CellPhone string `json:"cell_phone,omitempty"`
}

// Contains the attendee’s various different addresses. All are optional
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-addresses
type AttendeeAddresses []AttendeeAddress
type AttendeeAddress struct {
	// The attendee’s home address
	Home *Address `json:"home,omitempty"`
	// The attendee’s ship address
	Ship *Address `json:"ship,omitempty"`
	// The attendee’s workl address
	Work *Address `json:"work,omitempty"`
}

// A list of objects with answers to custom questions
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-answers

type AttendeeAnswers []AttendeeAnswer

type AttendeeAnswer struct {
	// The ID of the custom question
	QuestionID string `json:"question_id,omitempty"`
	// The text of the custom question
	Question string `json:"question,omitempty"`
	// One of multiple_choice, or text
	Type string `json:"type,omitempty"`
	// The attendee’s answer
	Answer string `json:"answer,omitempty"`
}

// A list of objects representing the barcodes for this order (usually only one per attendee)
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-barcodes
type AttendeeBarcodes []AttendeeBarcode
type AttendeeBarcode struct {
	//  The barcode contents. Note that if the event organizer has turned off printable
	// tickets, this field will be null in order to prevent exposing the barcode value
	Barcode string `json:"barcode,omitempty"`
	// One of unused, used, or refunded
	Status string `json:"status,omitempty"`
	// When the attendee barcode was created
	Created DateTime `json:"created,omitempty"`
	// When the attendee barcode was changed
	Changed DateTime `json:"changed,omitempty"`
}

// Represents team information for the attendee if the event has teams configured
//
// https://www.eventbrite.com/developer/v3/response_formats/attendee/#ebapi-attendee-team
type AttendeeTeam struct {
	// The team’s ID
	ID string `json:"id,omitempty"`
	// The team’s name
	Name string `json:"name,omitempty"`
	// When the attendee joined the team
	DateJoined DateTime `json:"date_joined,omitempty"`
	// The event the team is part of
	EventID string `json:"event_id,omitempty"`
}

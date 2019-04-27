package telegroid

type Update struct {
	UpdateId           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
	Poll               *Poll               `json:"poll"`
}

type WebhookInfo struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int      `json:"last_error_date"`
	LastErrorMessage     *string  `json:"last_error_message"`
	MaxConnections       int      `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

type User struct {
	Id           int     `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name"`
	Username     *string `json:"username"`
	LanguageCode *string `json:"language_code"`
}

type Chat struct {
	Id                          int        `json:"id"`
	Type                        ChatType   `json:"type"`
	Title                       *string    `json:"title"`
	Username                    *string    `json:"username"`
	FirstName                   *string    `json:"first_name"`
	LastName                    *string    `json:"last_name"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators"`
	Photo                       *ChatPhoto `json:"photo"`
	Description                 *string    `json:"description"`
	InviteLink                  *string    `json:"invite_link"`
	PinnedMessage               *Message   `json:"pinned_message"`
	StickerSetName              *string    `json:"sticker_set_name"`
	CanSetStickerName           bool       `json:"can_set_sticker_name"`
}

type Message struct {
	MessageId             int                `json:"message_id"`
	From                  *User              `json:"from"`
	Date                  int                `json:"date"`
	Chat                  *Chat              `json:"chat"`
	ForwardFrom           *User              `json:"forward_from"`
	ForwardFromChat       *Chat              `json:"forward_from_chat"`
	ForwardFromMessageId  int                `json:"forward_from_message_id"`
	ForwardSignature      *string            `json:"forward_signature"`
	ForwardSenderName     *string            `json:"forward_sender_name"`
	ForwardDate           int                `json:"forward_date"`
	ReplyToMessage        *Message           `json:"reply_to_message"`
	EditDate              int                `json:"edit_date"`
	MediaGroupId          *string            `json:"media_group_id"`
	AuthorSignature       *string            `json:"author_signature"`
	Text                  *string            `json:"text"`
	Entities              []MessageEntity    `json:"entities"`
	CaptionEntities       []MessageEntity    `json:"caption_entities"`
	Audio                 *Audio             `json:"audio"`
	Document              *Document          `json:"document"`
	Game                  *Game              `json:"game"`
	Photo                 []PhotoSize        `json:"photo"`
	Sticker               *Sticker           `json:"sticker"`
	Video                 *Video             `json:"video"`
	Voice                 *Voice             `json:"voice"`
	VideoNote             *VideoNote         `json:"video_note"`
	Caption               *string            `json:"caption"`
	Location              *Location          `json:"location"`
	Venue                 *Venue             `json:"venue"`
	NewChatMembers        []User             `json:"new_chat_members"`
	LeftChatMember        *User              `json:"left_chat_member"`
	NewChatTitle          *string            `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize        `json:"new_chat_photo"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`
	GroupChatCreated      bool               `json:"group_chat_created"`
	SupergroupChatCreated bool               `json:"supergroup_chat_created"`
	ChannelChatCreated    bool               `json:"channel_chat_created"`
	MigrateToChatId       int                `json:"migrate_to_chat_id"`
	MigrateFromChatId     int                `json:"migrate_from_chat_id"`
	PinnedMessage         *Message           `json:"pinned_message"`
	Invoice               *Invoice           `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
	ConnectedWebsite      *string            `json:"connected_website"`
	PassportData          *PassportData      `json:"passport_data"`
}

type MessageEntity struct {
	Type   EntityType `json:"type"`
	Offset int        `json:"offset"`
	Length int        `json:"length"`
	Url    *string    `json:"url"`
	User   *User      `json:"user"`
}

type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type MimeFile struct {
	FileId   string  `json:"file_id"`
	MimeType *string `json:"mime_type"`
	FileSize int     `json:"file_size"`
}

type Audio struct {
	MimeFile
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
}

type Document struct {
	MimeFile
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
}

type Video struct {
	MimeFile
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
}

type Voice struct {
	MimeFile
	Duration string `json:"duration"`
}

type VideoNote struct {
	FileId   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int        `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareId string   `json:"foursquare_id"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type Poll struct {
	Id       string       `json:"id"`
	Question string       `json:"question"`
	Options  []PollOption `json:"options"`
	IsClosed bool         `json:"is_closed"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileId   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageId *string  `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            *string  `json:"data"`
	GameShortName   *string  `json:"game_short_name"`
}

type ChatPhoto struct {
	SmallFileId string `json:"small_file_id"`
	BigFileId   string `json:"big_file_id"`
}

type ChatMember struct {
	User                  User         `json:"user"`
	Status                MemberStatus `json:"status"`
	UntilDate             int          `json:"until_date"`
	IsMember              bool         `json:"is_member"`
	CanBeEdited           bool         `json:"can_be_edited"`
	CanChangeInfo         bool         `json:"can_change_info"`
	CanPostMessages       bool         `json:"can_post_messages"`
	CanEditMessages       bool         `json:"can_edit_messages"`
	CanDeleteMessages     bool         `json:"can_delete_messages"`
	CanPinMessages        bool         `json:"can_pin_messages"`
	CanInviteUsers        bool         `json:"can_invite_users"`
	CanRestrictMembers    bool         `json:"can_restrict_members"`
	CanPromoteMembers     bool         `json:"can_promote_members"`
	CanSendMessages       bool         `json:"can_send_messages"`
	CanSendMediaMessages  bool         `json:"can_send_media_messages"`
	CanSendOtherMessages  bool         `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool         `json:"can_add_web_page_previews"`
}

type Sticker struct {
	FileId       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        *string       `json:"emoji"`
	SetName      *string       `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type StickerSet struct {
	Name         string    `json:"name"`
	Title        string    `json:"title"`
	ContainsMask bool      `json:"contains_mask"`
	Stickers     []Sticker `json:"stickers"`
}

type Animation struct {
	MimeFile
	Thumb    *PhotoSize `json:"thumb"`
	FileName *string    `json:"file_name"`
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

type CallbackGame struct{}

type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}

type InlineQuery struct {
	Id       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageId string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type SuccessfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int       `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionId        string    `json:"shipping_option_id"`
	OrderInfo               OrderInfo `json:"order_info"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id"`
}

type ShippingQuery struct {
	Id              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	Id               string    `json:"id"`
	From             User      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int       `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionId string    `json:"shipping_option_id"`
	OrderInfo        OrderInfo `json:"order_info"`
}

type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type EncryptedPassportElement struct {
	Type        PassportType    `json:"type"`
	Data        string          `json:"data"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Files       []*PassportFile `json:"files"`
	FrontSide   *PassportFile   `json:"front_side"`
	ReverseSide *PassportFile   `json:"reverse_side"`
	Selfie      *PassportFile   `json:"selfie"`
	Translation []*PassportFile `json:"translation"`
	Hash        string          `json:"hash"`
}

type PassportFile struct {
	FileId   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
	FileDate int64  `json:"file_date"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

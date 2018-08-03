package model

// User is user
type User struct {
	IsBusiness      bool `json:"is_business"`
	HasPlacedOrders bool `json:"has_placed_orders"`
	Nametag         struct {
		Mode          int    `json:"mode"`
		Gradient      int    `json:"gradient"`
		Emoji         string `json:"emoji"`
		SelfieSticker int    `json:"selfie_sticker"`
	} `json:"nametag"`
	CanSeeOrganicInsights      bool   `json:"can_see_organic_insights"`
	ShowInsightsTerms          bool   `json:"show_insights_terms"`
	Username                   string `json:"username"`
	FullName                   string `json:"full_name"`
	Pk                         int64  `json:"pk"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	ProfilePicID               string `json:"profile_pic_id"`
	CanBoostPost               bool   `json:"can_boost_post"`
	AllowedCommenterType       string `json:"allowed_commenter_type"`
	ProfilePicURL              string `json:"profile_pic_url"`
	IsVerified                 bool   `json:"is_verified"`
	ReelAutoArchive            string `json:"reel_auto_archive"`
	IsPrivate                  bool   `json:"is_private"`
	AllowContactsSync          bool   `json:"allow_contacts_sync"`
	PhoneNumber                string `json:"phone_number"`
}

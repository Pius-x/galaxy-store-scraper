package galaxy_store_scraper

type SamsungApp struct {
	CommentList []struct {
		ModifyDate        string `json:"modifyDate"`
		CreateDate        string `json:"createDate"`
		UserId            any    `json:"userId"`
		CommentText       string `json:"commentText"`
		CommentStatudCode any    `json:"commentStatudCode"`
		SellerAnswerFlag  string `json:"sellerAnswerFlag"`
		RegisterTypeCode  string `json:"registerTypeCode"`
		RatingValueNumber string `json:"ratingValueNumber"`
		LoginId           string `json:"loginId"`
	} `json:"commentList"`
	DetailMain struct {
		AppId                  any      `json:"appId"`
		ContentId              string   `json:"contentId"`
		ContentName            string   `json:"contentName"`
		ContentType            string   `json:"contentType"`
		ShopId                 string   `json:"shopId"`
		StoreType              any      `json:"storeType"`
		ContentDescription     string   `json:"contentDescription"`
		ContentNewDescription  string   `json:"contentNewDescription"`
		CnvrnImgUrl            string   `json:"cnvrnImgUrl"`
		YoutubeUrl             string   `json:"youtubeUrl"`
		YoutubeImgUrl          any      `json:"youtubeImgUrl"`
		ItemGroupId            string   `json:"itemGroupId"`
		ItemPurchaseFlag       string   `json:"itemPurchaseFlag"`
		PrepostFlag            string   `json:"prepostFlag"`
		SellerName             string   `json:"sellerName"`
		CopyrightHolder        string   `json:"copyrightHolder"`
		SellerId               any      `json:"sellerId"`
		IconURL                string   `json:"iconURL"`
		LocalPrice             string   `json:"localPrice"`
		DiscountPrice          any      `json:"discountPrice"`
		DiscountFlag           string   `json:"discountFlag"`
		FreeFlag               string   `json:"freeFlag"`
		CtntLanguageCode       string   `json:"ctntLanguageCode"`
		LimitAgeCd             string   `json:"limitAgeCd"`
		LimitAgeDetail         string   `json:"limitAgeDetail"`
		ContentGradeImageUrl   any      `json:"contentGradeImageUrl"`
		RatingValue            string   `json:"ratingValue"`
		RatingNumber           string   `json:"ratingNumber"`
		CommentListTotalCount  any      `json:"commentListTotalCount"`
		ContentBinaryVersion   string   `json:"contentBinaryVersion"`
		ContentBinarySize      string   `json:"contentBinarySize"`
		CustomerSupportEmail   string   `json:"customerSupportEmail"`
		SellerPrivatePolicy    string   `json:"sellerPrivatePolicy"`
		SellerSite             any      `json:"sellerSite"`
		CommentTitle           any      `json:"commentTitle"`
		EditorComment          any      `json:"editorComment"`
		CountryCode            string   `json:"countryCode"`
		Locale                 string   `json:"locale"`
		ModifyDate             string   `json:"modifyDate"`
		GeneralCategoryId      any      `json:"generalCategoryId"`
		GeneralCategoryName    any      `json:"generalCategoryName"`
		CategoryPath           any      `json:"categoryPath"`
		DeeplinkUrl            string   `json:"deeplinkUrl"`
		ManagerlinkUrl         string   `json:"managerlinkUrl"`
		PermissionList         []string `json:"permissionList"`
		ScreenshotRes          string   `json:"screenshotRes"`
		ContentBinaryWatchType string   `json:"contentBinaryWatchType"`
	} `json:"DetailMain"`
	PageMetaDesc     string `json:"page_meta_desc"`
	PageTitle        string `json:"page_title"`
	ContentId        string `json:"contentId"`
	UrlPath          string `json:"urlPath"`
	Platform         string `json:"platform"`
	BadgeUrlEndpoint string `json:"badgeUrlEndpoint"`
	Sticker          struct {
		PreviewImgUrl    any `json:"previewImgUrl"`
		StickerType      any `json:"stickerType"`
		SupportedAppList any `json:"supportedAppList"`
	} `json:"Sticker"`
	AppType    string `json:"appType"`
	AppId      string `json:"appId"`
	Screenshot struct {
		ScrnShtUrlList []struct {
			SmallScrnShtUrl    string `json:"smallScrnShtUrl"`
			OriginalScrnShtUrl string `json:"originalScrnShtUrl"`
		} `json:"scrnShtUrlList"`
		SmallScrnShtUrlList any `json:"smallScrnShtUrlList"`
		ScreenshotUrl       any `json:"screenshotUrl"`
		ScreenshotSeq       any `json:"screenshotSeq"`
		LangCd              any `json:"langCd"`
		ScrnShtHrztSize     any `json:"scrnShtHrztSize"`
		ScrnShtVrtcSize     any `json:"scrnShtVrtcSize"`
		ScreenshotCnt       any `json:"screenshotCnt"`
		ScreenshotRes       any `json:"screenshotRes"`
	} `json:"Screenshot"`
	SellerInfo struct {
		SellerType          string `json:"sellerType"`
		SellerTradeName     string `json:"sellerTradeName"`
		Representation      string `json:"representation"`
		SellerSite          any    `json:"sellerSite"`
		FirstName           string `json:"firstName"`
		LastName            any    `json:"lastName"`
		SellerNumber        string `json:"sellerNumber"`
		FirstSellerAddress  string `json:"firstSellerAddress"`
		SecondSellerAddress any    `json:"secondSellerAddress"`
		RegistrationNumber  any    `json:"registrationNumber"`
		ReportNumber        any    `json:"reportNumber"`
	} `json:"SellerInfo"`
	CommentListTotalCount string `json:"commentListTotalCount"`
}

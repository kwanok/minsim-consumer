package kafka

type RedditComment struct {
	Id            string  `json:"id"`
	Author        string  `json:"author"`
	Body          string  `json:"body"`
	BodyHtml      string  `json:"body_html"`
	CreatedUtc    int     `json:"created_utc"`
	Distinguished *string `json:"distinguished"`
	Edited        bool    `json:"edited"`
	IsSubmitter   bool    `json:"is_submitter"`
	LinkId        string  `json:"link_id"`
	ParentId      string  `json:"parent_id"`
	Permalink     string  `json:"permalink"`
	Stickied      bool    `json:"stickied"`
	Submission    string  `json:"submission"`
	Subreddit     string  `json:"subreddit"`
	SubredditId   string  `json:"subreddit_id"`
}

type RedditSubmission struct {
	Id                  string      `json:"id"`
	Author              string      `json:"author"`
	AuthorFlairText     string      `json:"author_flair_text"`
	Clicked             bool        `json:"clicked"`
	CreatedUtc          int         `json:"created_utc"`
	IsOriginalContent   bool        `json:"is_original_content"`
	IsSelf              bool        `json:"is_self"`
	LinkFlairTemplateId interface{} `json:"link_flair_template_id"`
	LinkFlairText       string      `json:"link_flair_text"`
	Locked              bool        `json:"locked"`
	Name                string      `json:"name"`
	NumComments         int         `json:"num_comments"`
	Over18              bool        `json:"over18"`
	Permalink           string      `json:"permalink"`
	Score               int         `json:"score"`
	Selftext            string      `json:"selftext"`
	Subreddit           string      `json:"subreddit"`
	SubredditId         string      `json:"subreddit_id"`
	Title               string      `json:"title"`
	UpvoteRatio         struct {
		Scale int    `json:"scale"`
		Value string `json:"value"`
	} `json:"upvote_ratio"`
	Url string `json:"url"`
}

package jsondemo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var url = "https://api.github.com/search/repositories?q=blog"

type JsonStruct struct {
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		ArchiveURL       string      `json:"archive_url"`
		Archived         bool        `json:"archived"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        string      `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		DeploymentsURL   string      `json:"deployments_url"`
		Description      string      `json:"description"`
		Disabled         bool        `json:"disabled"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            int         `json:"forks"`
		ForksCount       int         `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasProjects      bool        `json:"has_projects"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         string      `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HTMLURL          string      `json:"html_url"`
		ID               int         `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         interface{} `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		License          interface{} `json:"license"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NodeID           string      `json:"node_id"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       int         `json:"open_issues"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		Owner            struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int    `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"owner"`
		Permissions struct {
			Admin bool `json:"admin"`
			Pull  bool `json:"pull"`
			Push  bool `json:"push"`
		} `json:"permissions"`
		Private         bool    `json:"private"`
		PullsURL        string  `json:"pulls_url"`
		PushedAt        string  `json:"pushed_at"`
		ReleasesURL     string  `json:"releases_url"`
		Score           float64 `json:"score"`
		Size            int     `json:"size"`
		SSHURL          string  `json:"ssh_url"`
		StargazersCount int     `json:"stargazers_count"`
		StargazersURL   string  `json:"stargazers_url"`
		StatusesURL     string  `json:"statuses_url"`
		SubscribersURL  string  `json:"subscribers_url"`
		SubscriptionURL string  `json:"subscription_url"`
		SvnURL          string  `json:"svn_url"`
		TagsURL         string  `json:"tags_url"`
		TeamsURL        string  `json:"teams_url"`
		TreesURL        string  `json:"trees_url"`
		UpdatedAt       string  `json:"updated_at"`
		URL             string  `json:"url"`
		Watchers        int     `json:"watchers"`
		WatchersCount   int     `json:"watchers_count"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

func Run() {
	resp,err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var respResultJson JsonStruct
	err = json.NewDecoder(resp.Body).Decode(&respResultJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(respResultJson)
}

func Run2()  {
	var respResultJson JsonStruct
	//json.Unmarshal([]byte(data),&respResultJson)
	fmt.Println(respResultJson)
}

func Run3()  {
	params := make(map[string]interface{})

	params["name"] = "lecon"
	params["age"] = 111

	result,err := json.Marshal(params)
	if  err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(result))
}
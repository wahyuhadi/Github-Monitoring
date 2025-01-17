package main

import (
	"flag"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"time"
)


var (
	API = "https://api.github.com/search/code?q=isKey+user:isUser"
)

type Input struct {
	isUsername string
	isKeyword string
}

// REST API Models form Git documentation
type Response struct {
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		GitURL     string `json:"git_url"`
		HTMLURL    string `json:"html_url"`
		Name       string `json:"name"`
		Path       string `json:"path"`
		Repository struct {
			ArchiveURL       string      `json:"archive_url"`
			AssigneesURL     string      `json:"assignees_url"`
			BlobsURL         string      `json:"blobs_url"`
			BranchesURL      string      `json:"branches_url"`
			CollaboratorsURL string      `json:"collaborators_url"`
			CommentsURL      string      `json:"comments_url"`
			CommitsURL       string      `json:"commits_url"`
			CompareURL       string      `json:"compare_url"`
			ContentsURL      string      `json:"contents_url"`
			ContributorsURL  string      `json:"contributors_url"`
			DeploymentsURL   string      `json:"deployments_url"`
			Description      interface{} `json:"description"`
			DownloadsURL     string      `json:"downloads_url"`
			EventsURL        string      `json:"events_url"`
			Fork             bool        `json:"fork"`
			ForksURL         string      `json:"forks_url"`
			FullName         string      `json:"full_name"`
			GitCommitsURL    string      `json:"git_commits_url"`
			GitRefsURL       string      `json:"git_refs_url"`
			GitTagsURL       string      `json:"git_tags_url"`
			HooksURL         string      `json:"hooks_url"`
			HTMLURL          string      `json:"html_url"`
			ID               int         `json:"id"`
			IssueCommentURL  string      `json:"issue_comment_url"`
			IssueEventsURL   string      `json:"issue_events_url"`
			IssuesURL        string      `json:"issues_url"`
			KeysURL          string      `json:"keys_url"`
			LabelsURL        string      `json:"labels_url"`
			LanguagesURL     string      `json:"languages_url"`
			MergesURL        string      `json:"merges_url"`
			MilestonesURL    string      `json:"milestones_url"`
			Name             string      `json:"name"`
			NodeID           string      `json:"node_id"`
			NotificationsURL string      `json:"notifications_url"`
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
			Private         bool   `json:"private"`
			PullsURL        string `json:"pulls_url"`
			ReleasesURL     string `json:"releases_url"`
			StargazersURL   string `json:"stargazers_url"`
			StatusesURL     string `json:"statuses_url"`
			SubscribersURL  string `json:"subscribers_url"`
			SubscriptionURL string `json:"subscription_url"`
			TagsURL         string `json:"tags_url"`
			TeamsURL        string `json:"teams_url"`
			TreesURL        string `json:"trees_url"`
			URL             string `json:"url"`
		} `json:"repository"`
		Score float64 `json:"score"`
		Sha   string  `json:"sha"`
		URL   string  `json:"url"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

// Value receiver for flags
func (i Input) call() {
	if i.isUsername == "not" {
		fmt.Println("[!] Use -h for details")
		fmt.Println("[+] ./Git-Mon --user=wahyuhadi --keyword=password")
		return
	}
	rep := strings.NewReplacer("isKey", i.isKeyword,"isUser", i.isUsername)
	isURL := rep.Replace(API)
	res:= new(Response)
	getJson(isURL, res)
	if res.TotalCount == 0 {
		fmt.Println("[!] Nothing Found !!!! ")
		fmt.Println("[!] Total found for keyword " , i.isKeyword," is ", res.TotalCount)
		return
	}

	for _, element := range res.Items {
		
		fmt.Println("\n[+] URL   :",element.HTMLURL)
		fmt.Println("[+] Name  :",element.Name)
		fmt.Println("[+] Path  :",element.Path)
		fmt.Println("[+] Owner :",element.Repository.Owner.Login)
	}

	fmt.Println("\n[+] Finished ")
	fmt.Println("[+] Total Found is ", res.TotalCount)
	
}


// Get JSON before parsing

var myClient = &http.Client{Timeout: 10 * time.Second}
func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}

// Main Function
func main() {
	isUsername := flag.String("user", "not", "Github Account")
	isKeyword := flag.String("keyword", "", "Keyword will find")
	flag.Parse()
	isInput := Input{*isUsername, *isKeyword}
	isInput.call()
}
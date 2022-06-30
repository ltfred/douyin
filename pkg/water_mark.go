package pkg

import (
	"douyin-gui/constant"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"fyne.io/fyne/v2/widget"
)

func WaterMark(url string, path string, status *widget.Label) {
	urls := strings.Split(url, ",")
	if len(urls) == 1 && urls[0] == "" {
		status.SetText("Please enter the video url.")
		return
	}
	if path == "" {
		status.SetText("Please select folder.")
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	limit := make(chan struct{}, 10)
	for i := range urls {
		status.SetText("Downloading......")
		limit <- struct{}{}
		go func(url string) {
			defer wg.Done()
			Do(url, path)
			<-limit
		}(urls[i])
	}
	wg.Wait()
	status.SetText("Download complete.")
}

type Data struct {
	StatusCode int `json:"status_code"`
	ItemList   []struct {
		LabelTopText interface{} `json:"label_top_text"`
		Music        struct {
			Title       string `json:"title"`
			Author      string `json:"author"`
			CoverMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_medium"`
			Duration   int    `json:"duration"`
			ID         int64  `json:"id"`
			Mid        string `json:"mid"`
			CoverThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_thumb"`
			PlayURL struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_url"`
			Position interface{} `json:"position"`
			Status   int         `json:"status"`
			CoverHd  struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_hd"`
			CoverLarge struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_large"`
		} `json:"music"`
		RiskInfos struct {
			Type             int    `json:"type"`
			Content          string `json:"content"`
			ReflowUnplayable int    `json:"reflow_unplayable"`
			Warn             bool   `json:"warn"`
		} `json:"risk_infos"`
		Author struct {
			Nickname     string `json:"nickname"`
			Signature    string `json:"signature"`
			AvatarLarger struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_larger"`
			AvatarMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_medium"`
			UID         string `json:"uid"`
			AvatarThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			FollowersDetail  interface{} `json:"followers_detail"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			PolicyVersion    interface{} `json:"policy_version"`
			TypeLabel        interface{} `json:"type_label"`
			CardEntries      interface{} `json:"card_entries"`
			MixInfo          interface{} `json:"mix_info"`
			ShortID          string      `json:"short_id"`
			UniqueID         string      `json:"unique_id"`
			Geofencing       interface{} `json:"geofencing"`
		} `json:"author"`
		AwemeID    string      `json:"aweme_id"`
		Desc       string      `json:"desc"`
		Geofencing interface{} `json:"geofencing"`
		VideoText  interface{} `json:"video_text"`
		Promotions interface{} `json:"promotions"`
		GroupID    int64       `json:"group_id"`
		Images     interface{} `json:"images"`
		Statistics struct {
			AwemeID      string `json:"aweme_id"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
		} `json:"statistics"`
		VideoLabels interface{} `json:"video_labels"`
		CreateTime  int         `json:"create_time"`
		Duration    int         `json:"duration"`
		CommentList interface{} `json:"comment_list"`
		Video       struct {
			HasWatermark bool `json:"has_watermark"`
			Duration     int  `json:"duration"`
			PlayAddr     struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
			Cover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover"`
			Width       int `json:"width"`
			OriginCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"origin_cover"`
			Ratio        string `json:"ratio"`
			Height       int    `json:"height"`
			DynamicCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"dynamic_cover"`
			BitRate interface{} `json:"bit_rate"`
			Vid     string      `json:"vid"`
		} `json:"video"`
		ShareInfo struct {
			ShareWeiboDesc string `json:"share_weibo_desc"`
			ShareDesc      string `json:"share_desc"`
			ShareTitle     string `json:"share_title"`
		} `json:"share_info"`
		TextExtra []struct {
			HashtagID   int64  `json:"hashtag_id"`
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
		} `json:"text_extra"`
		ImageInfos interface{} `json:"image_infos"`
		ForwardID  string      `json:"forward_id"`
		GroupIDStr string      `json:"group_id_str"`
		ChaList    []struct {
			HashTagProfile string `json:"hash_tag_profile"`
			IsCommerce     bool   `json:"is_commerce"`
			Cid            string `json:"cid"`
			UserCount      int    `json:"user_count"`
			CoverItem      struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_item"`
			ViewCount    int         `json:"view_count"`
			ChaName      string      `json:"cha_name"`
			Desc         string      `json:"desc"`
			ConnectMusic interface{} `json:"connect_music"`
			Type         int         `json:"type"`
		} `json:"cha_list"`
		ShareURL     string      `json:"share_url"`
		IsLiveReplay bool        `json:"is_live_replay"`
		AwemeType    int         `json:"aweme_type"`
		LongVideo    interface{} `json:"long_video"`
		AuthorUserID int64       `json:"author_user_id"`
		IsPreview    int         `json:"is_preview"`
	} `json:"item_list"`
	FilterList []interface{} `json:"-"`
	Extra      struct {
		Logid string `json:"logid"`
		Now   int64  `json:"now"`
	} `json:"-"`
}

func DoRequest(videoID string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", constant.DYUrl, nil)
	request.Header.Set("user-agent", constant.UserAgent)
	q := request.URL.Query()
	q.Add("item_ids", videoID)
	request.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Request Error: %v", err)
		return nil, err
	}
	return resp, nil
}

func Do(url string, path string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request DouYin Video Url Error: %v", err)
		return
	}
	r := regexp.MustCompile("\\d+")
	all := r.FindAll([]byte(res.Request.URL.Path), 1)
	if len(all) == 0 {
		log.Fatal("Wrong Url")
		return
	}
	resp1, err := DoRequest(string(all[0]))
	if err != nil {
		return
	}
	defer resp1.Body.Close()
	body, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		log.Fatalf("Read Boby Error: %v", err)
		return
	}
	var result Data
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal Error: %v", err)
		return
	}
	// images
	if result.ItemList[0].Images != nil {
		parseImages(result.ItemList[0].Images, result.ItemList[0].Desc, path)
		return
	}
	// video
	parseVideo(result, path)
}

func parseImages(data interface{}, dir string, path string) {
	err := os.MkdirAll(path+"/"+dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Create Folder Error: %v", err)
	}
	d := data.([]interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(d))
	for i, v := range d {
		go func(v interface{}, i int) {
			defer wg.Done()
			urlList := v.(map[string]interface{})["url_list"]
			u := strings.TrimRight(fmt.Sprintf("%s", urlList), "]")
			u = strings.TrimLeft(u, "[")
			urls := strings.Split(u, " ")
			response, err := DoRequest(urls[0])
			if err != nil {
				return
			}
			defer response.Body.Close()
			fileName := fmt.Sprintf("%v/%v/tupian-%v", path, dir, i) + ".jpeg"
			file, _ := os.Create(fileName)
			_, _ = io.Copy(file, response.Body)
		}(v, i)
	}
	wg.Wait()
}

func parseVideo(result Data, path string) {
	videoUrlWm := result.ItemList[0].Video.PlayAddr.URLList[0]
	videoUrl := strings.Replace(videoUrlWm, "playwm", "play", 1)
	resp2, err := DoRequest(videoUrl)
	if err != nil {
		return
	}
	defer resp2.Body.Close()
	fileName := fmt.Sprintf("%v/%v", path, result.ItemList[0].Desc) + ".mp4"
	file, _ := os.Create(fileName)
	_, _ = io.Copy(file, resp2.Body)
}

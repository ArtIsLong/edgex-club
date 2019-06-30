package controller

import (
	"bytes"
	"edgex-club/cache"
	"edgex-club/model"
	"edgex-club/repository"
	"edgex-club/util"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	_ "time"

	mux "github.com/gorilla/mux"
)

type ReturnLoginUserToPageData struct {
	UserInfo    string
	Token       string
	UserPrePage string
}
type UserInfo struct {
	Id         int64
	Login      string
	Avatar_url string
}

func ValidToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	token := vars["token"]

	_, ok := cache.TokenCache[token]
	var isVaild string
	if ok {
		isVaild = "1" //有效
	} else {
		isVaild = "0" //无效
	}
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte(isVaild))
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	m := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	name := m["name"]
	pwd := m["password"]

	u := model.User{Name: name, Password: pwd}
	ok, err := repository.UserRepos.Exists(u)

	if err != nil {
		log.Println("User: " + name + " login failed : " + err.Error())
		w.Write([]byte("log failed : " + err.Error()))
		return
	}

	if ok {
		token := util.GetMd5String(name)
		cache.TokenCache[token] = u
		log.Println("User: " + name + " login.")
		m["avatarUrl"] = "https://avatars1.githubusercontent.com/u/42457890?v=4"
		m["id"] = "5bc0081dcedad5121dccebff"

		mm := make(map[string]interface{})
		mm["token"] = token
		mm["userInfo"] = m
		result, _ := json.Marshal(mm)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(result)
	} else {
		log.Println("User: " + name + " login failed : ")
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write([]byte("no user "))
	}

}

func LoginByGitHub(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := r.URL.Query()
	code := vars["code"][0]
	userPrePage := vars["state"][0]

	githubtoken := getGithubTokenByCode(code)
	githubUserInfo := getUserInfoByToken(githubtoken)

	var userInfo = UserInfo{}
	jsonStr := bytes.NewReader([]byte(githubUserInfo))
	json.NewDecoder(jsonStr).Decode(&userInfo)

	userName := userInfo.Login
	userId := strconv.FormatInt(userInfo.Id, 10)
	avatarUrl := userInfo.Avatar_url
	u := model.User{Name: userName, GitHubId: userId, AvatarUrl: avatarUrl}
	ok, _ := repository.UserRepos.ExistsByGitHub(u)
	if !ok {
		repository.UserRepos.Insert(u)
		log.Println("user not exist, it's a new user,save to db.")
	}
	u = repository.UserRepos.FindOneByName(u.Name)
	token := util.GetMd5String(u.Name)
	cache.TokenCache[token] = u

	log.Printf("User login: %v", u)
	//log.Println("User: " + u.Name + " login.")
	m := make(map[string]string)
	m["name"] = u.Name
	m["avatarUrl"] = u.AvatarUrl
	m["id"] = u.Id.Hex()
	mJson, err := json.Marshal(m)
	if err != nil {
		log.Printf("user : %v login failed !", u)
	}

	t, _ := template.ParseFiles("static/init.html")
	data := ReturnLoginUserToPageData{
		UserInfo:    string(mJson),
		Token:       token,
		UserPrePage: userPrePage,
	}

	t.Execute(w, data)
}

func getGithubTokenByCode(code string) string {
	url := "https://github.com/login/oauth/access_token"
	param := make(map[string]string, 10)
	param["client_id"] = "173d78b242d4fc35aca9"
	param["client_secret"] = "a35c510a599f19c6041325bcb7b3579072eb9228"
	param["code"] = code

	bytesData, err := json.Marshal(param)
	if err != nil {
		log.Println("param json faild!")
	}
	jsonStr := bytes.NewReader(bytesData)
	request, _ := http.NewRequest("POST", url, jsonStr)

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("request github to get code faild!")
	}
	// defer response.Body.Close()
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("respBytes  faild!")
	}
	//access_token=11f5681361f6448411abcb33401641fc61d9d6e4&scope=user&token_type=bearer
	tokenStr := string(respBytes)
	args := strings.Split(tokenStr, "&")
	tokenInfo := strings.Split(args[0], "=")
	token := tokenInfo[1]

	return token
}

func getUserInfoByToken(token string) string {
	url := "https://api.github.com/user?access_token=" + token + "&scope=user"
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	// defer resp.Body.Close()
	userData, _ := ioutil.ReadAll(resp.Body)
	userInfo := string(userData)
	//log.Println(userInfo)
	return userInfo
}

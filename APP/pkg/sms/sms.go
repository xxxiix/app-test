package sms

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"main/models"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func SmsResult(p *models.ParamSignUp, IsTest bool) (code string, err error) {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := viper.GetString("sms.account")   //用户名是登录用户中心->验证码短信->产品总览->APIID
	_password := viper.GetString("sms.password") //查看密码请登录用户中心->验证码短信->产品总览->APIKEY
	_mobile := p.Phone                           //获取手机号
	code = random()                              //获取随机验证码
	_content := "您的验证码是：" + code + "。请不要把验证码泄露给其他人。"
	if IsTest {
		return
	}
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	//body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	body := strings.NewReader(v.Encode()) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	var resp *http.Response
	resp, err = client.Do(req) //发送
	if err != nil {
		return
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	return
}

func random() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

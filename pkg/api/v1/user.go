package v1

import (
	"math"
	"strconv"
	"strings"
)

// User as user
type User struct {
	username  string
	Useragent string
	session   string
}

type login struct {
	DeviceID         string `json:"device_id"`
	LoginAttemptUser int    `json:"login_attempt_user"`
	Password         string `json:"password"`
	Username         string `json:"username"`
	Csrf             string `json:"_csrftoken"`
	UuID             string `json:"_uuid"`
}

var release = []string{
	"4.0.4", "4.3.1", "4.4.4", "5.1.1", "6.0.1",
}

var dpi = []string{
	"801", "577", "576", "538", "515", "424", "401", "373",
}

var resolution = []string{
	"3840x2160", "1440x2560", "2560x1440", "1440x2560",
	"2560x1440", "1080x1920", "1080x1920", "1080x1920",
}

func (u *User) info(username string) (string, string, string) {
	md5int, _ := u.api(username)

	info := Device[md5int%int64(len(Device))]

	var manufacturer = info[0]
	var device = info[1]
	var model = info[2]

	return manufacturer, device, model
}

func (u *User) userAgent(username string) string {
	var i, _ = u.api(username)

	var release = release[i%5]
	var dpi = dpi[i%8]
	var resolution = resolution[i%8]

	var manufacturer, device, model = u.info(username)

	var ua = []string{strconv.FormatInt(i, 10) + "/" + release, dpi + "dpi",
		resolution, manufacturer,
		model, device, "en-US"}

	agent := strings.Join(ua, "; ")

	return "Instagram " + AppVersion + " Android(" + agent + ")"
}

func (u *User) api(username string) (int64, string) {
	var hash = CreateMD5(username)
	var n, _ = strconv.ParseInt(hash, 10, 32)

	var f = int64(math.Round(float64(n) / 10e32))

	return f, string(18 + (f % 5))
}

// Login create user instance
func Login(username string, password string) *User {
	user := &User{}

	user.username = username
	user.Useragent = user.userAgent(username)

	http := Http{}

	body := GenerateSignature(login{
		DeviceID:         GenerateDeviceID(),
		LoginAttemptUser: 0,
		Password:         password,
		Username:         user.username,
		Csrf:             "missing",
		UuID:             GenerateUUID(),
	})

	http.
		Post("accounts/login/",
			"ig_sig_key_version=4&signed_body="+body).
		With(map[string]string{
			"User-Agent": user.Useragent,
		}).
		Exec()

	return user
}

// Using user session
func Using(session string) *User {
	user := &User{}

	user.session = session

	return user
}
